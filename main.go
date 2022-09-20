package main

import (
	"fmt"

	"github.com/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("joohwan")
	fmt.Println(account)
}
