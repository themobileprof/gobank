package bank

import (
	"errors"
)

// Create an interface with a Statement() string function.
type Bank interface {
	Statement() string
}

// Customer struct ...
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

// Welcome placeholder function
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

// Transfer ...
func (a *Account) Transfer(to *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to transfer should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("insufficient balance to transfer")
	}

	a.Balance -= amount
	to.Balance += amount
	return nil
}

// Statement ...
func Statement(b Bank) string {
	return string(b.Statement())
}
