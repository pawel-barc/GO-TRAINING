package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Print("No possibility to divide with 0")
		return 0
	}
	return a / b
}

func main() {
	var a, b float64
	var op string

	fmt.Println("Welcome in the calculator")
	fmt.Println("Enter first number")
	fmt.Scanln(&a)

	fmt.Println("Select one operation: (+, -, *, /) ")
	fmt.Scanln(&op)

	fmt.Println("Select second number")
	fmt.Scanln(&b)

	var result float64

	switch op {
	case "+":
		result = add(a, b)

	case "-":
		result = subtract(a, b)

	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		fmt.Println("Operation unknown!")
		return
	}
	fmt.Printf("Result : %.4f\n", result)
}
