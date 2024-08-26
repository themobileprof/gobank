package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/themobileprof/bank"
	"github.com/themobileprof/db"
)

func main() {
	db.Connect()
	var accounts = bank.Account{}

	fmt.Println(bank.Welcome())
	account, err := createAccount(&accounts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account successfully created for user: %v\n", account.Name)

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
