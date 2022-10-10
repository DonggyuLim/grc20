package token

import (
	"bytes"
	"fmt"
	"sync"

<<<<<<< HEAD
	"cosmossdk.io/math"
=======
	"github.com/shopspring/decimal"
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
)

type Token struct {
	Name        string
	Symbol      string
	Decimal     uint8
<<<<<<< HEAD
	TotalSupply math.Int
	Balance     map[string]math.Int
	Allowances  map[string]math.Int
	mutex       sync.RWMutex
=======
	TotalSupply decimal.Decimal
	Balance     map[string]decimal.Decimal
	Allowances  map[string]decimal.Decimal
	mutex       sync.Mutex
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
}

// Query
func (t *Token) GetName() string {
	return t.Name
}
func (t *Token) GetSymbol() string {
	return t.Symbol
}
<<<<<<< HEAD
func (t *Token) GetTotalSupply() math.Int {
	return t.TotalSupply
}

func (t *Token) BalanceOf(account string) math.Int {

	balance := t.Balance[account]
	if balance.IsNil() {
		return math.ZeroInt()
	}
	return balance
=======
func (t *Token) GetTotalSupply() decimal.Decimal {
	return t.TotalSupply
}

func (t *Token) BalanceOf(account string) decimal.Decimal {
	return t.Balance[account]
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
}

func (t *Token) GetDecimal() uint8 {
	return t.Decimal
}

<<<<<<< HEAD
func (t *Token) Allowance(owner, spender string) math.Int {
	return t.allowance(owner, spender)
}
func (t *Token) allowance(owner, spender string) math.Int {
=======
func (t *Token) Allowance(owner, spender string) decimal.Decimal {
	return t.allowance(owner, spender)
}
func (t *Token) allowance(owner, spender string) decimal.Decimal {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
	key := owner + ":" + spender
	if t.Allowances[key].IsNil() {
		return math.ZeroInt()
	}
	return t.Allowances[key]
}

// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// Mutate

<<<<<<< HEAD
func (t *Token) Transfer(from, to string, amount math.Int) error {
=======
func (t *Token) Transfer(from, to string, amount decimal.Decimal) error {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
	t.mutex.Lock()
	defer t.mutex.Unlock()
	err := t.checkBalance(from, amount)
	if err != nil {
		return err
	}
	t.transfer(from, to, amount)
	return nil
}
<<<<<<< HEAD

func (t *Token) transfer(from, to string, amount math.Int) {
	t.Balance[from] = t.BalanceOf(from).Sub(amount)
	t.Balance[to] = t.BalanceOf(to).Add(amount)
}

func (t *Token) Approve(owner, spender string, amount math.Int) error {
=======
func (t *Token) transfer(from, to string, amount decimal.Decimal) {

	fromBalance := t.Balance[from]
	t.Balance[from] = fromBalance.Sub(amount)
	toBalance := t.Balance[to]
	t.Balance[to] = toBalance.Add(amount)
}

func (t *Token) Approve(owner, spender string, amount decimal.Decimal) error {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if err := t.checkBalance(owner, amount); err != nil {
		return err
	}
	t.approve(owner, spender, amount)
	return nil
}

<<<<<<< HEAD
func (t *Token) approve(owner, spender string, amount math.Int) error {
=======
func (t *Token) approve(owner, spender string, amount decimal.Decimal) error {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)

	key := owner + ":" + spender
	currentAllowance := t.Allowances[key]

<<<<<<< HEAD
	t.Allowances[key] = currentAllowance.Add(amount)
=======
	t.Allowances[key] = currentBalance.Add(amount)
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
	t.Balance[owner] = t.BalanceOf(owner).Sub(amount)

	return nil
}

<<<<<<< HEAD
func (t *Token) TransferFrom(owner, spender, to string, amount math.Int) error {
=======
func (t *Token) TransferFrom(owner, spender, to string, amount decimal.Decimal) error {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
	t.mutex.Lock()
	defer t.mutex.Unlock()
	//allownace 가 amount 보다 작은지 확인
	if err := t.checkspendAllowance(owner, spender, amount); err != nil {
		return err
	}
	t.transferfrom(owner, spender, amount)
	return nil
}

<<<<<<< HEAD
func (t *Token) transferfrom(owner, spender string, amount math.Int) {
	key := owner + ":" + spender
	t.Allowances[key] = t.Allowance(key, spender).Sub(amount)

	if t.Allowance(owner, spender).Equal(math.ZeroInt()) {
=======
func (t *Token) transferfrom(owner, spender string, amount decimal.Decimal) {
	key := owner + ":" + spender
	t.Allowances[key] = t.allowance(owner, spender).Sub(amount)
	zero := decimal.NewFromInt(0)
	if t.Allowances[key].Cmp(zero) != 1 {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
		delete(t.Allowances, key)
	}
	t.Balance[spender] = t.BalanceOf(spender).Add(amount)
}

<<<<<<< HEAD
func (t *Token) Mint(account string, amount math.Int) {
=======
func (t *Token) Mint(account string, amount decimal.Decimal) {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.mint(account, amount)
}

<<<<<<< HEAD
func (t *Token) mint(account string, amount math.Int) {

	t.TotalSupply = t.TotalSupply.Add(amount)
	t.Balance[account] = t.BalanceOf(account).Add(amount)

}

func (t *Token) Burn(address string, amount math.Int) {
=======
func (t *Token) mint(address string, amount decimal.Decimal) {

	t.TotalSupply = t.TotalSupply.Add(amount)
	currentBalance := t.BalanceOf(address)
	// newBalance := currentBalance + amount
	newBalance := currentBalance.Add(amount)
	t.Balance[address] = newBalance
}

func (t *Token) Burn(address string, amount decimal.Decimal) {
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.burn(address, amount)
}

<<<<<<< HEAD
func (t *Token) burn(address string, amount math.Int) {
	currentBalance := t.BalanceOf(address)
	newBalance := currentBalance.Sub(amount)
	t.TotalSupply = t.GetTotalSupply().Sub(amount)
	if newBalance.LT(math.ZeroInt()) {
		t.Balance[address] = math.ZeroInt()
	} else {
=======
func (t *Token) burn(address string, amount decimal.Decimal) {
	currentBalance := t.BalanceOf(address)
	newBalance := currentBalance.Sub(amount)
	zero := decimal.NewFromInt(0)
	if newBalance.Cmp(zero) == -1 {
		t.TotalSupply = t.GetTotalSupply().Sub(amount)
		t.Balance[address] = zero
	} else {
		t.TotalSupply = t.GetTotalSupply().Sub(amount)
>>>>>>> parent of 3f9e7f6 (modified number Deciaml.Decimal -> uint64)
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
