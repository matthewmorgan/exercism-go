package account

type Account struct {
	balance int64
	closed  bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	a := Account{initialDeposit, false}
	return &a
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.closed {
		return 0, false
	}
	if amount < 0 && -1*amount > a.balance {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
