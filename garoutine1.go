package main

import (
	"fmt"
	"time"
)

func addMoney(balance *int, amount int) {
	for i := 0; i < 5; i++ {
		*balance += amount
		fmt.Println("Added ", amount, "-> balance now: ", *balance)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	balance := 100

	go addMoney(&balance, 10)
	go addMoney(&balance, 20)

	time.Sleep(2 * time.Second)
	fmt.Println("Final balance:", balance)

}