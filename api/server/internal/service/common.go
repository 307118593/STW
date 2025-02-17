package service

import (
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

// ConvBalance 根据mult值判断进行乘法或除法运算，并保留两位小数
func ConvBalance(amount *big.Int, mult int64) string {
	// 将 amount 转换为 decimal
	amountDec := decimal.NewFromBigInt(amount, 0)
	// 计算 10 的幂的指数
	factor := decimal.NewFromFloat(math.Pow(10, float64(math.Abs(float64(mult)))))

	var result decimal.Decimal
	if mult < 0 {
		// 负数：除以 10 的幂
		result = amountDec.Div(factor)
	} else {
		// 正数：乘以 10 的幂
		result = amountDec.Mul(factor)
	}

	// 保留两位小数 并避免科学计数
	result = result.Round(8)

	// 返回浮动类型的结果
	return result.String()
}

// StringToWei 将 Ether（float64 类型）转换为 Wei（big.Int 类型），使用 decimal 包
func StringToWei(amount string) (*decimal.Decimal, error) {
	// 使用 decimal.NewFromFloat 将 float64 转换为 decimal 类型
	amountDec, _ := decimal.NewFromString(amount)

	// 1 Ether = 10^18 Wei
	wei := decimal.NewFromInt(1000000000000000000) // 10^18 Wei

	// 计算金额在 Wei 中的表示
	amountInWei := amountDec.Mul(wei)

	return &amountInWei, nil
}
