package Interface

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"cosmossdk.io/math"
	"github.com/DonggyuLim/grc20/db"
	"github.com/shopspring/decimal"
)

type GRC20 interface {
	GetName() string
	GetSymbol() string
<<<<<<< HEAD
	GetTotalSupply() math.Int
	BalanceOf(account string) math.Int
	GetDecimal() uint8
	Transfer(from, to string, amount math.Int) error
	TransferFrom(from, to, sepnder string, amount math.Int) error
	Allowance(owner, spender string) math.Int
	Approve(owner, spender string, amount math.Int) error
	Mint(account string, amount math.Int)
=======
	GetTotalSupply() decimal.Decimal
	BalanceOf(account string) decimal.Decimal
	GetDecimal() uint8
	Transfer(from, to string, amount decimal.Decimal) error
	TransferFrom(from, to, sepnder string, amount decimal.Decimal) error
	Allowance(owner, spender string) decimal.Decimal
	Approve(owner, spender string, amount decimal.Decimal) error
	Mint(account string, amount decimal.Decimal)
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
}

// GRC20struct -> byte
func StructToByte(data GRC20) []byte {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)

	err := enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println(result.Bytes())
	return result.Bytes()
}

// save Token data
func SaveToken(name string, token GRC20) {
	db.Add(name, StructToByte(token))
}
