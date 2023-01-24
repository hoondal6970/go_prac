package main

import (
	"fmt"
	"log"

	"github.com/hoondal6970/learngo/banking"
)

func main() {
	account := banking.NewAccount("solda")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(120)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance(), account.Owner())
}
