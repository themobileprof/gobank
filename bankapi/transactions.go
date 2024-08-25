package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/themobileprof/bank"
	"github.com/themobileprof/db"
)

func deposit(w http.ResponseWriter, req *http.Request) {
	account := &bank.Account{}
	account_id := 0

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

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		row := db.DB.QueryRow("SELECT u.name, a.id, a.balance FROM users u JOIN accounts a ON u.id = a.user_id WHERE account_number=?", numberqs)
		if err := row.Scan(&account.Name, &account_id, &account.Balance); err != nil {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			// Update account struct
			err := account.Deposit(amount)

			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				// Update the deposit in the database
				updateDeposit(w, account_id, account)

				// Print the statement
				fmt.Fprint(w, account.Statement())
			}
		}
	}
}

func updateDeposit(w http.ResponseWriter, account_id int, account *bank.Account) {
	// Update the deposit in the database
	_, err := db.DB.Exec("UPDATE accounts SET balance = ? WHERE id = ?", account.Balance, account_id)
	if err != nil {
		fmt.Fprintf(w, "UpdateDeposit: %v", err)
	} else {
		fmt.Fprintf(w, "Deposit updated successfully!")
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
			fmt.Fprint(w, account.Statement())
		}
	}
}
