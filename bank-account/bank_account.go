package account

import (
	"sync"
)

// BankAccount represents a bank account
type BankAccount interface {
	Balance() (balance int, ok bool)
	Close() (payout int, ok bool)
	Deposit(amount int) (balance int, ok bool)
}

// Open a bank account
func Open(initialBalance int) BankAccount {

	if initialBalance < 0 {
		return nil
	}

	return &bankAccountState{
		balance:  initialBalance,
		isActive: true,
	}
}

// Balance returns guess what? Balance of a bank account, of course
func (bas *bankAccountState) Balance() (balance int, ok bool) {

	return bas.runTransaction(func() (int, bool) {
		return bas.balance, true
	})
}

// Close closes bank account
func (bas *bankAccountState) Close() (payout int, ok bool) {
	return bas.runTransaction(func() (int, bool) {
		p := bas.balance

		bas.isActive = false
		bas.balance = 0

		return p, true
	})
}

// Deposit allows to deposit and withdraw (with negative amount) money from a bank account
func (bas *bankAccountState) Deposit(amount int) (balance int, ok bool) {
	return bas.runTransaction(func() (int, bool) {

		if amount < 0 && -1*amount > bas.balance {
			return bas.balance, false
		}

		bas.balance += amount
		return bas.balance, true
	})
}

type bankAccountState struct {
	balance  int
	isActive bool
	lock     sync.Mutex
}

func (bas *bankAccountState) runTransaction(f func() (int, bool)) (result int, ok bool) {
	bas.lock.Lock()
	defer bas.lock.Unlock()

	if !bas.isActive {
		return 0, false
	}

	return f()
}
