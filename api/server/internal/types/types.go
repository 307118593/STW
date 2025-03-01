// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package types

type BlockChainListHandlerResponse struct {
	BlockChainList []string `json:"blockChainList"`
}

type CreateWalletHandlerRequest struct {
	Blockchain string `json:"backchain"`
}

type CreateWalletHandlerResponse struct {
	Address string `json:"address"`
}

type Defi struct {
	Name     string `json:"name"`
	Amount   string `json:"amount"`
	Contract string `json:"contract"`
}

type ImportWalletHandlerRequest struct {
	Blockchain string `json:"backchain"`
	PrivateKey string `json:"privateKey"`
}

type ImportWalletHandlerResponse struct {
	Address string `json:"address"`
}

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type Transaction struct {
	Hash  string `json:"hash"`
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	Fee   string `json:"fee"`
	Time  string `json:"time"`
}

type WalletDefiTransferHandlerRequest struct {
	Blockchain string `json:"backchain"`
	PrivateKey string `json:"privateKey"`
	Contract   string `json:"contract"`
	From       string `json:"from"`
	To         string `json:"to"`
	Value      string `json:"value"`
	Fee        string `json:"fee"`
	Name       string `json:"name,optional"`
}

type WalletDefiTransferHandlerResponse struct {
	Hash string `json:"hash"`
}

type WalletInfoRequest struct {
	Blockchain string `json:"backchain"`
	Address    string `json:"address"`
}

type WalletInfoResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Defi    []Defi `json:"Defi"`
}

type WalletTransactionDetailHandlerRequest struct {
	Hash string `json:"hash"`
}

type WalletTransactionDetailHandlerResponse struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	Fee   string `json:"fee"`
	Time  string `json:"time"`
}

type WalletTransactionHandlerRequest struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type WalletTransactionHandlerResponse struct {
	TransacteList []Transaction `json:"transacteList"`
}

type WalletTransferHandlerRequest struct {
	Blockchain string `json:"backchain"`
	PrivateKey string `json:"privateKey"`
	From       string `json:"from"`
	To         string `json:"to"`
	Value      string `json:"value"`
	Fee        string `json:"fee"`
}

type WalletTransferHandlerResponse struct {
	Hash string `json:"hash"`
}
