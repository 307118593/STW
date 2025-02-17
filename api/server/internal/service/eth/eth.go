package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"sync"
	"time"
	"wallet/api/server/internal/service"
	"wallet/api/server/internal/service/eth/abis"
	"wallet/api/server/internal/types"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial("https://sepolia.drpc.org")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to Ethereum client")

}

func NewBlockNumber() (uint64, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(header.Number.String()) // 5671744
	return header.Number.Uint64(), nil
}

func CreateWallet() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err.Error())
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//address := hexutil.Encode(publicKeyBytes)[4:]
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address, privateKeyString
}

// 导入钱包
func ImportWallet(privateKey string) (string, error) {
	//privateKeyBytes, err := hexutil.Decode("0x" + privateKey)
	//if err != nil {
	//	return "", err
	//}

	// 解析私钥
	private, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}

	// 从私钥获取公钥
	publicKey := private.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA format")
	}

	// 通过公钥获取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}
	return address, nil
}

// address : 0xDA6B8652F84235D6634Dea70aA9CE6c6EC041Fd0
var contracts = map[string]string{
	"STW":        "0xD6048051aCfFc51B8B1E12fcc45F65652Bf018e5",
	"RareTron":   "0x54FA517F05e11Ffa87f4b22AE87d91Cec0C2D7E1",
	"VanityTron": "0x489c5CB7fD158B0A9E7975076D758268a756C025",
	"stw2":       "0x35c3E414C12dfB9184E7D0843C5801D8f388108D",
}

func WalletInfo(address string) (address1 string, balance string, difi []types.Defi) {
	//查询eth余额
	address1 = address
	account := common.HexToAddress(address)
	balanceMain, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err.Error())
	}
	balance = service.ConvBalance(balanceMain, -18)
	//balance = service.ConvBalance(big.Int{}, -18)

	//使用协程查询代币余额,协程控制2s,协程内利用channel返回结果
	difi = make([]types.Defi, 0)
	ch := make(chan types.Defi, len(contracts)) // 创建一个带缓冲区的 channel
	var wg sync.WaitGroup
	// 并发执行查询操作
	for name, contract := range contracts {
		wg.Add(1)
		go func(name string, contract string, address string) {
			defer wg.Done()
			// 设置超时时间为 2 秒
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			// 创建一个 channel 用于返回查询结果
			resultCh := make(chan types.Defi)

			// 使用 select 监听超时和查询结果
			go func() {
				symbol, bal := getDifiBalance(address, contract)

				// 模拟某个协程超时（例如: name == "stw2"）
				if name == "stw2" {
					// 手动控制某个协程超时
					//time.Sleep(5 * time.Second)
				}

				// 将查询结果发送到 resultCh
				resultCh <- types.Defi{Name: symbol, Amount: bal, Contract: contract}
			}()

			select {
			case <-ctx.Done(): // 如果超时
				if ctx.Err() == context.DeadlineExceeded {
					log.Printf("Timeout for contract %s\n", name)
					return
				}
			case result := <-resultCh:
				// 如果查询成功，发送结果到主 channel
				ch <- result
			}
		}(name, contract, address)
	}

	// 等待所有协程完成
	wg.Wait()
	close(ch)
	// 从 channel 中接收所有查询结果并存入 difi 切片
	for result := range ch {
		difi = append(difi, result)
	}
	return
}

// 查询代币余额
func getDifiBalance(address string, contractAddress string) (symbol string, balance string) {
	tokenAddress := common.HexToAddress(contractAddress)
	instance, err := abis.NewToken(tokenAddress, client)
	if err != nil {
		panic(err.Error())
	}

	addressToken := common.HexToAddress(address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, addressToken)
	if err != nil {
		panic(err.Error())
	}

	symbol, err = instance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err.Error())
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("contractAddress", contractAddress)
	fmt.Println("contractSymbol", symbol)
	fmt.Println("contractDecimals", decimals)
	fmt.Println("contractBal", bal)

	//balance = service.ConvBalance(bal.Int64(), -18)
	balance = service.ConvBalance(bal, -int64(decimals))
	return
}

// eth转账.暂不考虑并发问题
func WalletTransfer(address string, privateKey string, to string, value string, fee string) (hash string) {
	private, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err.Error())
	}
	fromAddress := common.HexToAddress(address)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("nonce:", nonce)
	valueD, _ := service.StringToWei(value)
	valueBigInt := valueD.BigInt()
	fmt.Println("valueBigInt:", valueBigInt)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	fmt.Println("gasPrice:", gasPrice)
	if err != nil {
		panic(err.Error())
	}

	toAddress := common.HexToAddress(to)

	var data []byte
	tx := ethTypes.NewTransaction(nonce, toAddress, valueBigInt, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err.Error())
	}

	signedTx, err := ethTypes.SignTx(tx, ethTypes.NewEIP155Signer(chainID), private)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("signedTx", signedTx.Hash().Hex())
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		//panic(err.Error())
		fmt.Println("err", err)
		panic(err.Error())
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	hash = signedTx.Hash().Hex()
	return
}

func DefiTransfer(address string, privateKey string, to string, value string, fee string, contract string) (hash string) {
	private, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err.Error())
	}

	publicKey := private.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err.Error())
	}

	ethValue := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err.Error())
	}

	toAddress := common.HexToAddress(to)
	tokenAddress := common.HexToAddress(contract)

	transferFnSignature := []byte("transfer(address,uint256)")
	//hash := sha3.NewKeccak256()
	//202502 sha3包被移动到crypto
	hash1 := crypto.NewKeccakState()
	hash1.Write(transferFnSignature)
	methodID := hash1.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	//amount := new(big.Int)
	//amount.SetString("1000000000000000000000", 10) // 1000 tokens
	valueD, _ := service.StringToWei(value)
	amount := valueD.BigInt()
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(gasLimit) // 23256

	tx := ethTypes.NewTransaction(nonce, tokenAddress, ethValue, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err.Error())
	}

	signedTx, err := ethTypes.SignTx(tx, ethTypes.NewEIP155Signer(chainID), private)
	if err != nil {
		panic(err.Error())
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err.Error())
	}
	hash = signedTx.Hash().Hex()
	return
}
