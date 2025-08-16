package main

import (
	"fmt"
)

var (
	currentBalance float64
)

func BankServer(transactions <-chan Transaction) {
	for tx := range transactions {
		switch tx.action {
		case "deposit":
			currentBalance += tx.amount
		case "withdraw":
			currentBalance -= tx.amount
		case "checkBalance":
			tx.reply <- fmt.Sprintf("Current Balance %.1f", currentBalance)
		}
	}
}

type Transaction struct {
	action string
	amount float64
	reply  chan string
}

func main() {
	transactions := []Transaction{
		{action: "deposit", amount: 1000, reply: make(chan string)},
		{action: "checkBalance", reply: make(chan string)},
		{action: "deposit", amount: 2000, reply: make(chan string)},
		{action: "checkBalance", reply: make(chan string)},
		{action: "withdraw", amount: 2000, reply: make(chan string)},
		{action: "checkBalance", reply: make(chan string)}}

	ch := make(chan Transaction, len(transactions))

	go BankServer(ch)

	for _, tx := range transactions {
		ch <- tx
		if tx.action == "checkBalance" {
			fmt.Println(<-tx.reply)
		}
	}

	close(ch)
}
