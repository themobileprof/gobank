package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/themobileprof/bank"
	"github.com/themobileprof/db"
	"golang.org/x/exp/rand"
)

// createAccount creates a new bank account and inserts it into the database.
// It generates a random 10-digit account number that starts with 001.
// The account details such as customer name, email, phone, address, gender, and date of birth are set to default values.
// The created account is then inserted into the database using the insertAccount function.
// If the insertion is successful, the created account is returned along with nil error.
// If there is an error during the insertion, an empty account and the corresponding error message are returned.
func createAccount(accounts *bank.Account) (*bank.Account, error) {

	// Generate a random 10 digits account number that starts with 001
	rand.Seed(uint64(time.Now().UnixNano()))
	accountNumber := "001" + strconv.Itoa(rand.Intn(9000000)+1000000)

	*accounts = bank.Account{
		Customer: bank.Customer{
			Name:    "John Doe",
			Email:   "john@gmail.com",
			Phone:   "(213) 555 0147",
			Address: "Los Angeles, California",
			Gender:  "Male",
			DoB:     "1983-10-10",
		},
		Number: accountNumber,
	}

	_, err := insertAccount(accounts)
	if err != nil {
		return &bank.Account{}, fmt.Errorf("insertAccountError: %v", err)
	}

	return accounts, nil
}

// insertAccount inserts the given account into the database.
// It checks if the account is nil and returns an error if it is.
// It also checks if the database connection is nil and returns an error if it is.
// The account details such as name, email, phone number, address, gender, and date of birth are inserted into the "users" table.
// The account number and balance are inserted into the "accounts" table.
// If the insertion is successful, the ID of the inserted account is returned along with nil error.
// If there is an error during the insertion, the corresponding error message is returned.
func insertAccount(accounts *bank.Account) (int64, error) {
	if accounts == nil {
		return 0, fmt.Errorf("account is nil")
	}

	// Assuming db is a global variable or passed as a parameter
	if db.DB == nil {
		return 0, fmt.Errorf("database connection is nil")
	}

	result, err := db.DB.Exec("INSERT INTO users (name, email, phone_number, address, gender, date_of_birth) VALUES (?, ?, ?, ?, ?, ?)", accounts.Name, accounts.Email, accounts.Phone, accounts.Address, accounts.Gender, accounts.DoB)
	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}
	user_id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}

	result, err = db.DB.Exec("INSERT INTO accounts (user_id, account_number, balance) VALUES (?, ?, ?)", user_id, accounts.Number, accounts.Balance)
	if err != nil {
		return 0, fmt.Errorf("addAccount: %v", err)
	}

	account_id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAccount: %v", err)
	}
	return account_id, nil
}

// getAccountByNumber retrieves the account with the given account number from the database.
// It checks if the database connection is nil and returns an error if it is.
// It queries the "users" and "accounts" tables to retrieve the account details.
// If the account is found, its details are assigned to the account variable and returned along with nil error.
// If the account is not found, an error message is returned.
func getAccountByNumber(number string) (*bank.Account, error) {
	account := &bank.Account{}

	if db.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	row := db.DB.QueryRow("SELECT u.name, u.email, u.phone_number, u.address, u.gender, u.date_of_birth, a.account_number, a.balance FROM users u JOIN accounts a ON u.id = a.user_id WHERE a.account_number = ?", number)
	if err := row.Scan(&account.Name, &account.Email, &account.Phone, &account.Address, &account.Gender, &account.DoB, &account.Number, &account.Balance); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("account not found")
		}
		return nil, fmt.Errorf("getAccountByNumber %v: %v", number, err)
	}
	return account, nil
}
