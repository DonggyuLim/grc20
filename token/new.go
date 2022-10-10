package token

<<<<<<< HEAD
import (
	"cosmossdk.io/math"
)

func NewToken(name, symbol string, dec uint8) *Token {
	bm := make(map[string]math.Int)
	am := make(map[string]math.Int)
=======
import "github.com/shopspring/decimal"

func NewToken(name, symbol string, dec uint8) *Token {
	bm := make(map[string]decimal.Decimal)
	am := make(map[string]decimal.Decimal)
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)

	t := Token{
		Name:        name,
		Symbol:      symbol,
		Decimal:     dec,
<<<<<<< HEAD
		TotalSupply: math.ZeroInt(),
=======
		TotalSupply: decimal.NewFromInt(0),
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
		Balance:     bm,
		Allowances:  am,
	}
	return &t
}
