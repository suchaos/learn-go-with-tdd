package pointers

import (
	"errors"
	"fmt"
)

type Bitcode int

func (b Bitcode) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcode
}

/*
注意：接受者的类型为指针
*/
func (w *Wallet) Deposit(amount Bitcode) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcode {
	return w.balance
}

var ErrInsufficientFounds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(bitcode Bitcode) error {
	if w.balance < bitcode {
		return ErrInsufficientFounds
	}

	w.balance -= bitcode
	return nil
}
