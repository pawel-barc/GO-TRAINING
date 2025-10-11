package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("How many grades you have: ")
	fmt.Scanln(&n)

	grades := make([]float64, n)
	var sum float64

	for _, i := 0; i < n; i++ {
		fmt.Println("Enter a grade", i+1)
		fmt.Scanln(&grades[i])
		sum += grades[i]

	}
	average := sum / float64(n)

	fmt.Println("Here are your grades: ")
	for grade := range grades {
		fmt.Printf("\n%.2f", grade)
	}
	fmt.Printf("\nYour grades average is: %.2f\n", average)

	if average < 3 {
		fmt.Println("\nYou are an empty loser")
	} else if average < 4.5 {
		fmt.Println("Not bad bustard")
	} else if average <= 6 {
		fmt.Print("\nSmart Cookie")
	} else {
		fmt.Println("Something went wrong")
	}
}
