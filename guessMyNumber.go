package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1
	var guess int
	attempts := 0

	fmt.Println("Welcome to the 'Guess the number' game!")
	fmt.Println("I'm thinking of a number between 1 and 100")

	for {
		fmt.Println("Enter your guess")
		fmt.Scanln(&guess)
		attempts++
		if guess < secret {
			fmt.Println("To low, try again!")
		} else if guess > secret {
			fmt.Println("To high, try again!")
		} else {
			fmt.Printf("Bingo you guessed it in %d attempts.\n ", attempts)
			break
		}
	}
}
