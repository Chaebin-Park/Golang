package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (theAccount *Account) Deposit(amount int) {
	theAccount.balance += amount
}

var noMoney = "Can't withdraw"
var errNoMoney = errors.New(noMoney)

// Balance of the account
func (theAccount Account) Balance() int {
	return theAccount.balance
}

// WithDraw x amount from your account
func (theAccount *Account) WithDraw(amount int) error {
	if theAccount.balance < amount {
		return errNoMoney
	}
	theAccount.balance -= amount
	return nil
}

// ChangeOwner of the account
func (theAccount *Account) ChangeOwner(newOwner string) {
	theAccount.owner = newOwner
}

// Owner of the account
func (theAccount Account) Owner() string {
	return theAccount.owner
}

// toString 재정의 같은거
func (theAccount Account) String() string {
	return fmt.Sprint(theAccount.owner, "'s account.\nHas: ", theAccount.balance)
}