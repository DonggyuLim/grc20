package token

import (
	"bytes"
	"fmt"
	"sync"

	"cosmossdk.io/math"
)

type Token struct {
	Name        string
	Symbol      string
	Decimal     uint8
	TotalSupply math.Int
	Balance     map[string]math.Int
	Allowances  map[string]math.Int
	mutex       sync.RWMutex
}

// Query
func (t *Token) GetName() string {
	return t.Name
}
func (t *Token) GetSymbol() string {
	return t.Symbol
}
func (t *Token) GetTotalSupply() math.Int {
	return t.TotalSupply
}

func (t *Token) BalanceOf(account string) math.Int {
	return t.Balance[account]
}

func (t *Token) GetDecimal() uint8 {
	return t.Decimal
}

func (t *Token) Allowance(owner, spender string) math.Int {
	return t.allowance(owner, spender)
}
func (t *Token) allowance(owner, spender string) math.Int {
	key := owner + ":" + spender

	return t.Allowances[key]
}

// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// Mutate

func (t *Token) Transfer(from, to string, amount math.Int) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	err := t.checkBalance(from, amount)
	if err != nil {
		return err
	}
	t.transfer(from, to, amount)
	return nil
}
func (t *Token) transfer(from, to string, amount math.Int) {

	fromBalance := t.Balance[from]
	t.Balance[from] = fromBalance.Sub(amount)
	toBalance := t.Balance[to]
	t.Balance[to] = toBalance.Add(amount)
}

func (t *Token) Approve(owner, spender string, amount math.Int) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if err := t.checkBalance(owner, amount); err != nil {
		return err
	}
	t.approve(owner, spender, amount)
	return nil
}

func (t *Token) approve(owner, spender string, amount math.Int) error {

	key := owner + ":" + spender
	currentAllowance := t.Allowances[key]

	t.Allowances[key] = currentAllowance.Add(amount)
	t.Balance[owner] = t.BalanceOf(owner).Sub(amount)

	return nil
}

func (t *Token) TransferFrom(owner, spender, to string, amount math.Int) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	//allownace ??? amount ?????? ????????? ??????
	if err := t.checkspendAllowance(owner, spender, amount); err != nil {
		return err
	}
	t.transferfrom(owner, spender, amount)
	return nil
}

func (t *Token) transferfrom(owner, spender string, amount math.Int) {
	key := owner + ":" + spender
	t.Allowances[key] = t.allowance(owner, spender).Sub(amount)

	if t.Allowances[key].LTE(math.ZeroInt()) {
		delete(t.Allowances, key)
	}
	t.Balance[spender] = t.BalanceOf(spender).Add(amount)
}

func (t *Token) Mint(account string, amount math.Int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.mint(account, amount)
}

func (t *Token) mint(address string, amount math.Int) {

	t.TotalSupply = t.TotalSupply.Add(amount)
	currentBalance := t.BalanceOf(address)
	// newBalance := currentBalance + amount
	newBalance := currentBalance.Add(amount)
	t.Balance[address] = newBalance
}

func (t *Token) Burn(address string, amount math.Int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.burn(address, amount)
}

func (t *Token) burn(address string, amount math.Int) {
	currentBalance := t.BalanceOf(address)
	newBalance := currentBalance.Sub(amount)

	if newBalance.LT(math.ZeroInt()) {
		t.TotalSupply = t.GetTotalSupply().Sub(amount)
		delete(t.Balance, address)
	} else {
		t.TotalSupply = t.GetTotalSupply().Sub(amount)
		t.Balance[address] = newBalance
	}
}

func (t *Token) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, t.Name, t.Symbol, t.Allowances, t.Balance, t.Decimal, &t.mutex)
	return b.Bytes(), nil
}

func (t *Token) UnmarshalBinary(data []byte) error {

	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &t.Name, &t.Symbol, &t.Allowances, &t.Balance, &t.Decimal, &t.mutex)
	return err
}

// ///////////////////////////////////
// //////////////////////////////////
