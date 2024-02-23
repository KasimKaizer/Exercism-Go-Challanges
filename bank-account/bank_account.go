// Package account contains solution for the Bank Account exercise on Exercism.
package account

import "sync"

// Account type defines an account in the bank.
type Account struct {
	Closed bool
	Amount int64
	MX     sync.Mutex
}

// Open takes an amount and returns an account where input amount is its balance.
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{Amount: amount}
}

// Balance method returns the current balance of the account.
func (a *Account) Balance() (int64, bool) {
	a.MX.Lock()
	defer a.MX.Unlock() // lock any changes to account until this function is done executing.
	if a.Closed {
		return a.Amount, false
	}
	return a.Amount, true
}

// Deposit method allows amount to be deposited to or withdrawn from the account.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.MX.Lock()
	defer a.MX.Unlock()
	if a.Closed || (a.Amount+amount) < 0 {
		// if the account is closed or if the input amount would lead to a negative balance then
		// return false.
		return a.Amount, false
	}
	a.Amount += amount
	return a.Amount, true

}

// Close method closes the current account
func (a *Account) Close() (int64, bool) {
	a.MX.Lock()
	defer a.MX.Unlock()
	if a.Closed {
		return a.Amount, false
	}
	bal := a.Amount
	a.Amount, a.Closed = 0, true
	return bal, true
}
