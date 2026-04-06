package pointersanderrs

import (
	"errors"
	"fmt"
)

//Used as a single source of truth for our error message.
// It is a var, so it is available to the package (pointersanderrs)
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// we can create new types from existing ones.
// this creates a type Bitcoin, that is just an int.
type Bitcoin int
func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct{
	balance Bitcoin
}

//*** NOTE: We are passing a pointer to the receiver to maintain state of properties in 'this'
//*** "pointer to wallet"
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin{
	return w.balance
}
func(w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance{
		return ErrInsufficientFunds
	}
	
	w.balance -= amount
	return nil
}