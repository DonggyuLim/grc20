package Interface

import (
	"bytes"
	"encoding/gob"
	"log"

	"cosmossdk.io/math"
	"github.com/DonggyuLim/grc20/db"
)

type GRC20 interface {
	GetName() string
	GetSymbol() string
	GetTotalSupply() math.Int
	BalanceOf(account string) math.Int
	GetDecimal() uint8
	Transfer(from, to string, amount math.Int) error
	TransferFrom(from, to, sepnder string, amount math.Int) error
	Allowance(owner, spender string) math.Int
	Approve(owner, spender string, amount math.Int) error
	Mint(account string, amount math.Int)
}

// GRC20struct -> byte
func StructToByte(data GRC20) []byte {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)

	err := enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}

	return result.Bytes()
}

// save Token data
func SaveToken(name string, token GRC20) {
	db.Add(name, StructToByte(token))
}
