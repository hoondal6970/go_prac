package banking

import "errors"

//BankAccount struct
type Account struct {
	owner   string
	balance int
}

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit add the deposit and updates the value of balance
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance of your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of your account
func (a Account) Owner() string {
	return a.owner
}
