package token

import (
	"sync"
)

type Token struct {
	Name        string
	Symbol      string
	Decimal     uint8
	TotalSupply uint64
	Balance     map[string]uint64
	Allowances  map[string]uint64
	mutex       sync.Mutex
}

// Query
func (t *Token) GetName() string {
	return t.Name
}
func (t *Token) GetSymbol() string {
	return t.Symbol
}
func (t *Token) GetTotalSupply() uint64 {
	return t.TotalSupply
}

func (t *Token) BalanceOf(account string) uint64 {
	return t.Balance[account]
}

func (t *Token) GetDecimal() uint8 {
	return t.Decimal
}

func (t *Token) Allowance(owner, spender string) uint64 {
	return t.allowance(owner, spender)
}
func (t *Token) allowance(owner, spender string) uint64 {
	key := owner + ":" + spender
	return t.Allowances[key]
}

// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// Mutate

func (t *Token) Transfer(from, to string, amount uint64) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	err := t.checkBalance(from, amount)
	if err != nil {
		return err
	}
	t.transfer(from, to, amount)
	return nil
}

func (t *Token) transfer(from, to string, amount uint64) {

	fromBalance := t.Balance[from]
	t.Balance[from] = fromBalance - amount
	toBalance := t.Balance[to]
	t.Balance[to] = toBalance + amount
}

func (t *Token) Approve(owner, spender string, amount uint64) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if err := t.checkBalance(owner, amount); err != nil {
		return err
	}
	t.approve(owner, spender, amount)
	return nil
}

func (t *Token) approve(owner, spender string, amount uint64) error {

	key := owner + ":" + spender
	currentBalance := t.Allowances[key]

	t.Allowances[key] = currentBalance + amount
	t.Balance[owner] = t.BalanceOf(owner) - amount

	return nil
}

func (t *Token) TransferFrom(owner, spender, to string, amount uint64) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	//allownace 가 amount 보다 작은지 확인
	if err := t.checkspendAllowance(owner, spender, amount); err != nil {
		return err
	}
	t.transferfrom(owner, spender, amount)
	return nil
}

func (t *Token) transferfrom(owner, spender string, amount uint64) {
	key := owner + ":" + spender
	t.Allowances[key] = t.allowance(owner, spender) - amount

	if t.Allowances[key] == 0 {
		delete(t.Allowances, key)
	}
	t.Balance[spender] = t.BalanceOf(spender) + amount
}

func (t *Token) Mint(account string, amount uint64) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.mint(account, amount)
}

func (t *Token) mint(address string, amount uint64) {

	t.TotalSupply = t.TotalSupply + amount
	t.Balance[address] = t.BalanceOf(address) + amount
}

func (t *Token) Burn(address string, amount uint64) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.burn(address, amount)
}

func (t *Token) burn(address string, amount uint64) {
	currentBalance := t.BalanceOf(address)
	newBalance := int64(currentBalance) - int64(amount)

	if newBalance < 0 {
		t.TotalSupply = t.GetTotalSupply() - amount
		t.Balance[address] = 0
	} else {
		t.TotalSupply = t.GetTotalSupply() - amount
		t.Balance[address] = uint64(newBalance)
	}
}

// ///////////////////////////////////
// //////////////////////////////////
