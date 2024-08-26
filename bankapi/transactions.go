package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/themobileprof/bank"
	"github.com/themobileprof/db"
)

type accountStatement struct {
	Name    string
	Address string `json:"Address,omitempty"`
	Phone   string
	Number  string `json:"Account Number"`
	Balance float64
}

func (s accountStatement) Statement() string {
	return fmt.Sprintf("{Name: %v} {Address: %v} {Phone: %v} {Account Number: %v} {Balance: %v}", s.Name, s.Address, s.Phone, s.Number, s.Balance)
}

func deposit(w http.ResponseWriter, req *http.Request) {

	if db.DB == nil {
		fmt.Fprintf(w, "database connection is nil")
		return
	}

	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if _, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, err := getAccountByNumber(numberqs)
		if err != nil {
			fmt.Fprintf(w, "Error getting account: %v", err)
		} else {
			// Update account struct
			err := account.Deposit(amount)

			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				// Synchronize with database
				updateBalance(w, account)

				// Print the statement
				statement := accountStatement{
					Name:    account.Name,
					Number:  account.Number,
					Balance: account.Balance,
				}
				fmt.Fprint(w, account.Statement(statement))
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if _, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, err := getAccountByNumber(numberqs)
		if err != nil {
			fmt.Fprintf(w, "Error getting account: %v", err)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				// Synchronize with database
				updateBalance(w, account)

				// Print the statement
				statement := accountStatement{
					Name:    account.Name,
					Address: account.Address,
					Phone:   account.Phone,
					Number:  account.Number,
					Balance: account.Balance,
				}
				fmt.Fprint(w, account.Statement(statement))
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	fromqs := req.URL.Query().Get("from")
	toqs := req.URL.Query().Get("to")
	amountqs := req.URL.Query().Get("amount")

	if fromqs == "" || toqs == "" {
		fmt.Fprintf(w, "You need two account numbers to complete a transfer!")
		return
	}

	if _, err := strconv.ParseFloat(fromqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid debiting account number!")
	} else if _, err := strconv.ParseFloat(toqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid receiving account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Amount is invalid!")
	} else {
		fromAccount, err := getAccountByNumber(fromqs)
		if err != nil {
			fmt.Fprintf(w, "Error getting Debit account: %v", err)
		} else {
			toAccount, err := getAccountByNumber(toqs)
			if err != nil {
				fmt.Fprintf(w, "Error getting Receiving account: %v", err)
			} else {
				err := fromAccount.Transfer(toAccount, amount)
				if err != nil {
					fmt.Fprintf(w, "%v", err)
				} else {
					// Synchronize with database
					updateBalance(w, fromAccount)
					updateBalance(w, toAccount)

					// Print the statement
					statement := accountStatement{
						Name:    fromAccount.Name,
						Address: fromAccount.Address,
						Phone:   fromAccount.Phone,
						Number:  fromAccount.Number,
						Balance: fromAccount.Balance,
					}
					fmt.Fprint(w, fromAccount.Statement(statement))
				}
			}
		}
	}
}

func updateBalance(w http.ResponseWriter, account *bank.Account) bool {
	// Update the deposit in the database
	_, err := db.DB.Exec("UPDATE accounts SET balance = ? WHERE account_number = ?", account.Balance, account.Number)
	if err != nil {
		fmt.Fprintf(w, "UpdateDeposit: %v", err)
		return false
	} else {
		return true
	}

}

func statement(w http.ResponseWriter, req *http.Request) {

	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	// ensure the account number is like a number
	if _, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		// Get the account from the database

		account, err := getAccountByNumber(numberqs)
		if err != nil {
			fmt.Fprintf(w, "Error getting account: %v", err)
		} else {
			// Print the statement
			statement := accountStatement{
				Name:    account.Name,
				Address: account.Address,
				Phone:   account.Phone,
				Number:  account.Number,
				Balance: account.Balance,
			}
			fmt.Fprint(w, account.Statement(statement))
		}
	}
}
