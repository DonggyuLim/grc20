package token

import (
	"errors"
)

// check
func (t *Token) checkBalance(owner string, amount uint64) error {
	balance := t.BalanceOf(owner)
	if balance < amount {
		return errors.New("amount is biger than owner balance")
	}
	return nil
}

func (t *Token) checkspendAllowance(owner, spender string, amount uint64) error {
	allowance := t.allowance(owner, spender)
	if allowance < amount {
		return errors.New("amount is biger than allowance")
	}
	return nil
}
