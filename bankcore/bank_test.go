package bank

import (
	"strconv"
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Samuel",
			Address: "3 Thorborn Avenue, Sabo, Yaba, Lagos",
			Phone:   "(234) 803 395 4301",
		},
		Number:  "1001",
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  "1001",
		Balance: 0,
	}

	err := account.Deposit(10)
	if err != nil {
		t.Error("balance is not being updated after a deposit")
	}

	if account.Balance != 10 {
		t.Error("balance is not reflecting the right amount after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  "1001",
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil { // Note the err == nil for negative test
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  "1001",
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  "0012345267",
		Balance: 0,
	}

	// create a struct to hold a custom account
	// and then attach the Statement() method to it
	type customAccount struct {
		*Account
	}

	// create a function variable that implements the Statement() method
	statementFunc := func(c *customAccount) string {
		return c.Account.Number + " - " + c.Account.Customer.Name + " - " + strconv.FormatFloat(c.Account.Balance, 'f', -1, 64)
	}

	// create a custom account
	acc := &customAccount{&account}

	acc.Deposit(100)
	statement := statementFunc(acc)

	if statement != "0012345267 - John - 100" {
		t.Errorf("statement doesn't have the proper format: %v", statement)
	}
}
