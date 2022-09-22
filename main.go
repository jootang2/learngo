package main

import (
	"fmt"

	"github.com/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("joohwan")
	account.Deposit(10)
	fmt.Println(account)
}
