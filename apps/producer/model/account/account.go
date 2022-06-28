package account

import (
	"fmt"
	"github.com/atilasantos/go-banksimulation/apps/producer/model/person"
)

type Account struct {
	person  person.Person
	balance float64
}

func (a *Account) HasBalance(amount float64) bool {
	return a.balance >= amount
}

func (a *Account) Withdraw(amount float64) {
	if a.HasBalance(amount) {
		a.balance -= amount
	} else {
		fmt.Println("No balance available")
	}
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) TransferTo(destinationAccount Account, amount float64) {
	if a.HasBalance(amount) {
		a.Withdraw(amount)
		destinationAccount.Deposit(amount)
	} else {
		fmt.Println("No balance available")
	}

}
