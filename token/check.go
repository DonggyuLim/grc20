package token

import (
	"errors"

<<<<<<< HEAD
	"cosmossdk.io/math"
)

// check
func (t *Token) checkBalance(owner string, amount math.Int) error {
	balance := t.BalanceOf(owner)

	if balance.LT(amount) {
=======
	"github.com/shopspring/decimal"
)

// check
func (t *Token) checkBalance(owner string, amount decimal.Decimal) error {
	balance := t.BalanceOf(owner)
	if balance.Cmp(amount) == -1 {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
		return errors.New("amount is biger than owner balance")
	}
	return nil
}

<<<<<<< HEAD
func (t *Token) checkspendAllowance(owner, spender string, amount math.Int) error {
	allowance := t.Allowance(owner, spender)
	if allowance.LT(amount) {
=======
func (t *Token) checkspendAllowance(owner, spender string, amount decimal.Decimal) error {
	Allowance := t.allowance(owner, spender)
	if Allowance.Cmp(amount) == -1 {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
		return errors.New("amount is biger than allowance")
	}
	return nil
}
