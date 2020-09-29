package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}
func (w *Wallet) Deposit (amount Bitcoin)  {
	fmt.Printf("address of balance in deposit is %v \n", &w.balance)
	w.balance += amount
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("Poor ass")
	}

	w.balance -= amount
	return nil
}
func (w *Wallet) Balance() Bitcoin  {
	return w.balance
}
func (b Bitcoin) string() string  {
	return fmt.Sprintf("%d BTC", b)
}