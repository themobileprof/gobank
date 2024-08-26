package main

import (
	"testing"

	"github.com/themobileprof/bank"
	"github.com/themobileprof/db"
)

func TestCreateAccount(t *testing.T) {
	// Connect DB
	db.ConnectTesting()

	// Create a new account
	accounts := &bank.Account{}

	_, err := createAccount(accounts)
	if err != nil {
		t.Errorf("account not created. %s", err)
	}

	// Check if the account is created
	if accounts.Name != "John Doe" {
		t.Error("account name is incorrect")
	}

	// Check if the account number starts with 001
	if accounts.Number[:3] != "001" {
		t.Error("account number should start with 001")
	}

	// Check if the account has a balance of zero
	if accounts.Balance != 0.00 {
		t.Error("account balance is not zero")
	}
}

func TestInsertAccount(t *testing.T) {
	// Connect DB
	db.ConnectTesting()

	// Create a new account
	accounts := &bank.Account{
		Customer: bank.Customer{
			Name:    "Jane Doe",
			Email:   "jane@gmail.com",
			Phone:   "(213) 555 0147",
			Address: "Los Angeles, California",
			Gender:  "Female",
			DoB:     "2003-01-01",
		},
		Number:  "0017286376",
		Balance: 0.00,
	}

	// Insert the account
	_, err := insertAccount(accounts)
	if err != nil {
		t.Errorf("account not inserted. %s", err)
	}

	// Check if the account is created
	if accounts.Name != "Jane Doe" {
		t.Error("account not created")
	}

	// Check if the account number starts with 001
	if accounts.Number[:3] != "001" {
		t.Error("account number should start with 001")
	}

	// Check if the account has a balance of zero
	if accounts.Balance != 0.00 {
		t.Error("account balance is not zero")
	}
}

func TestGetAccountByNumber(t *testing.T) {
	// Connect DB
	db.ConnectTesting()

	// Create a new account
	accounts := &bank.Account{
		Customer: bank.Customer{
			Name:    "Sam Song",
			Email:   "samuel@gmail.com",
			Phone:   "(803) 555 0147",
			Address: "Lagos, Nigeria",
			Gender:  "Male",
			DoB:     "1983-10-10",
		},
		Number:  "0018989351",
		Balance: 0.00,
	}

	// Insert the account
	_, err := insertAccount(accounts)
	if err != nil {
		t.Errorf("account not inserted. %s", err)
	}

	// Get the account by number
	account, err := getAccountByNumber(accounts.Number)
	if err != nil {
		t.Errorf("failed to get account. %s", err)
	}

	// Check if the retrieved account matches the inserted account
	if account.Name != accounts.Name || account.Email != accounts.Email || account.Phone != accounts.Phone || account.Address != accounts.Address || account.Gender != accounts.Gender || account.DoB != accounts.DoB || account.Number != accounts.Number || account.Balance != accounts.Balance {
		t.Error("retrieved account does not match inserted account")
	}
}
