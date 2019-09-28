package account

import "sync"

type Account struct {
	balance int64
	closed  bool
	mux sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	a := Account{balance: initialDeposit, closed: false}
	return &a
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	// I used this pattern because there are multiple returns from these methods.
	defer a.mux.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.closed {
		return 0, false
	}
	if amount < 0 && -1*amount > a.balance {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
