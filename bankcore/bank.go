package bank

import (
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
	Name    string
	Email   string
	Phone   string
	Address string
	Gender  string
	DoB     string
}

// Account ...
type Account struct {
	Customer
	Number  string
	Balance float64
}

// Welcome
func Welcome() string {
	return "Welcome!"
}

// Deposit ...
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be less than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// Statement ...
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}
