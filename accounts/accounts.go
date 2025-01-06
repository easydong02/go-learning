package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errorNoMoney = errors.New("Can't withdraw. you are poor.")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 원본으로 사용
func (account *Account) Deposit(amount int) {
	account.balance += amount
}

// 복사본으로 사용
func (account Account) Balance() int {
	return account.balance
}

func (account Account) Owner() string {
	return account.owner
}

func (account *Account) Withdraw(amount int) error {
	if account.balance >= amount {
		account.balance -= amount
	} else {
		return errorNoMoney
	}

	return nil
}

func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

// String은 java의 toString과 같음
func (account Account) String() string {
	return fmt.Sprint(account.Owner(), "'s account.\nHas: ", account.Balance())
}
