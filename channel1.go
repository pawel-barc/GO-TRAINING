package main

import "fmt"

func deposit(balance *float64, amount float64, done chan bool) {
	*balance += amount
	fmt.Printf("Deposited %.2f -> balance now: %.2f (address: %p)\n", amount, *balance, balance)
	done <- true
}

func main() {
	var myBalance float64 = 100
	done := make(chan bool)

	fmt.Printf("Initial balance: %.2f (address: %p)\n", myBalance, &myBalance)

	go deposit(&myBalance, 50, done)
	go deposit(&myBalance, 30, done)

	<-done
	<-done

	fmt.Printf("Final balance: %.2f (address: %p)\n", myBalance, &myBalance)
}

