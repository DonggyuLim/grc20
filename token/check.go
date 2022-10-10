package token

import (
	"errors"

	"cosmossdk.io/math"
)

// check
func (t *Token) checkBalance(owner string, amount math.Int) error {
	balance := t.BalanceOf(owner)

	if balance.LT(amount) {
		return errors.New("amount is biger than owner balance")
	}
	return nil
}

func (t *Token) checkspendAllowance(owner, spender string, amount math.Int) error {
	allowance := t.Allowance(owner, spender)
	if allowance.LT(amount) {
		return errors.New("amount is biger than allowance")
	}
	return nil
}
