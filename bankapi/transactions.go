package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/themobileprof/bank"
	"github.com/themobileprof/db"
)

// accountStatement represents a bank account statement.
type accountStatement struct {
	Name    string  // Name of the account holder.
	Address string  `json:"Address,omitempty"` // Address of the account holder. (optional)
	Phone   string  // Phone number of the account holder.
	Number  string  `json:"Account Number"` // Account number of the bank account.
	Balance float64 // Current balance of the bank account.
}

// Statement returns the account statement as a JSON string.
// It marshals the accountStatement struct into JSON format and returns it as a string.
// If there is an error during the marshaling process, it returns the error message as a string.
func (a *accountStatement) Statement() string {
	json, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}

	return string(json)
}

// deposit is a function that handles the deposit operation for a bank account.
// It takes in an http.ResponseWriter and an http.Request as parameters.
// The function first checks if the database connection is nil, and if so, it returns an error message.
// It then retrieves the account number and amount from the query parameters of the request.
// If the account number is missing or invalid, an error message is returned.
// If the account number is valid, the function retrieves the account details from the database.
// It then updates the account balance by depositing the specified amount.
// If there is an error during the deposit operation, an error message is returned.
// Otherwise, the function synchronizes the updated account details with the database.
// Finally, it generates a statement for the account and returns it as a response.
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
				fmt.Fprint(w, statement.Statement())
			}
		}
	}
}

// withdraw is a function that handles the withdrawal of funds from a bank account.
// It takes in an http.ResponseWriter and an http.Request as parameters.
// The function retrieves the account number and withdrawal amount from the request URL query parameters.
// If the account number is missing, it returns an error message.
// If the account number is invalid, it returns an error message.
// If the withdrawal amount is invalid, it returns an error message.
// Otherwise, it retrieves the account information using the account number.
// If there is an error retrieving the account, it returns an error message.
// If the withdrawal is successful, it updates the account balance in the database.
// Finally, it generates a statement with the account information and returns it as a response.
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
				fmt.Fprint(w, statement.Statement())
			}
		}
	}
}

// transfer is a function that handles the transfer of funds between two accounts.
// It takes in the http.ResponseWriter and *http.Request as parameters.
// The function retrieves the "from", "to", and "amount" query parameters from the request URL.
// If either "from" or "to" is empty, it returns an error message indicating that two account numbers are required to complete a transfer.
// If the account numbers are valid, it retrieves the corresponding accounts using the getAccountByNumber function.
// If any error occurs during the retrieval of accounts, it returns an error message.
// Otherwise, it calls the Transfer method on the "fromAccount" to transfer the specified amount to the "toAccount".
// If an error occurs during the transfer, it returns the error message.
// After a successful transfer, it updates the balances of both accounts in the database using the updateBalance function.
// Finally, it generates a statement for the "fromAccount" and writes it to the http.ResponseWriter.
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
					fmt.Fprint(w, statement.Statement())
				}
			}
		}
	}
}

// updateBalance updates the balance of an account in the database.
// It takes a http.ResponseWriter and a *bank.Account as parameters.
// It returns a boolean value indicating whether the balance update was successful or not.
// If the update fails, it writes the error message to the http.ResponseWriter.
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

// statement is a handler function that generates and returns the account statement for a given account number.
// It expects the account number to be provided as a query parameter in the request URL.
// If the account number is missing or invalid, it returns an appropriate error message.
// If the account number is valid, it retrieves the account details from the database and generates the statement.
// The statement includes the account holder's name, address, phone number, account number, and balance.
// The generated statement is written to the http.ResponseWriter.
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
			fmt.Fprint(w, statement.Statement())
		}
	}
}
