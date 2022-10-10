package token

import (
	"cosmossdk.io/math"
)

func NewToken(name, symbol string, dec uint8) *Token {
	bm := make(map[string]math.Int)
	am := make(map[string]math.Int)

	t := Token{
		Name:        name,
		Symbol:      symbol,
		Decimal:     dec,
		TotalSupply: math.ZeroInt(),
		Balance:     bm,
		Allowances:  am,
	}
	return &t
}
