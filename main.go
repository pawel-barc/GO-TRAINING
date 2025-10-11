package main

import "fmt"

// func main() {
// 	name := "Paul"
// 	age := 25
// 	fmt.Printf("Hi, My name is %s and I'm %d years old", name, age)
// }

// func main() {
// 	fmt.Print("1) Print - inline \n")
// 	fmt.Println("2) Println - outline")
// 	fmt.Printf("3) Printf - %s is %d years old", "Paul", 40)
// }
// func main() {
// 	name := "Paul"
// 	age := "40"
// 	pi := 3.14
// 	fmt.Println(name, age, pi)
// }
// func main() {
// 	for i := 1; i <= 7; i++ {
// 	fmt.Println("number ", i)
// 	}

// }

// func main() {
// 	age := 18

// 	if age < 18 {
// 		fmt.Print("You are underage")
// 	}else {
// 		fmt.Print("You are of legal age")
// 	}
// }

// func main() {
// 	fruits := []string{"apple", "banana", "cherry"}

// 	for index, fruit := range fruits {
// 		fmt.Println(index, fruit)
// 	}
// }

// func greet(name string) string {
// 	return "Hallo, " + name + "!"
// }

// func main() {
// 	message := greet("Paul")
// 	fmt.Print(message)
// }

// func main() {
// 	a := 10.0
// 	b := 7.0

// 	fmt.Println("Addition: ", a+b )
// 	fmt.Println("Substraction: ", a-b )
// 	fmt.Println("Multiplication: ", a*b)
// 	fmt.Println("Division: ", a/b)
// }

// func main() {
// 	var birthYear int
// 	fmt.Print("Enter your birth year ")
// 	fmt.Scanln(&birthYear)

// 	currentYear := 2025
// 	fmt.Print("You are ", currentYear - birthYear)
// }

// func main() {
// 	var number int
// 	fmt.Print("Enter a number")
// 	fmt.Scanln(&number)

// 	if number%2 == 0 {
// 		fmt.Print("The number is even")
// 	} else {
// 		fmt.Print("Number is odd")
// 	}
// }

// func main() {
// 	var number int
// 	fmt.Print("Enter a number")
// 	fmt.Scanln(&number)
// 	sum := 0
// 	for i := 0; i <= number; i++ {
// 		sum += number
// 	}
// 	fmt.Print(sum)

// }

// func main () {
// 	sum := 0
// 	for i := 0; i <= 5; i++ {
// 		sum += i
// 	}
// 	fmt.Print("The sum is equal: ", sum )
// }

// func main () {
// 	names := []string{"Paul", "Alice", "Anne"}
// 	for i, name := range names {
// 		fmt.Println(i, name)
// 	}
// }

// func add (a int, b int) int {
// 	return a + b
// }
// func main() {
// 	result := add(5, 7)
// 	fmt.Print("Result is: ", result)
// }

// func main () {
// 	for i := 1; i <= 5; i++ {
// 		for j := 1; j <= 5; j++ {
// 			fmt.Print(i*j, "\t")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	text := "How many letters"
// 	fmt.Println("Number of letters: ", len(text))
// }

// func add (a int, b int) int {
// 	return a + b
// }

// func subst(a int, b int) int {
// 	return a - b
// }

// func mult(a int, b int ) int {
// 	return a*b
// }
// func divid(a int, b int) int {
// 	return a/b
// }

// func main() {
// 	var action string
// 	var a, b int
// 	fmt.Print("Enter two numbers and action: +,-,*,/ ")
// 	fmt.Scanln(&a, &b, &action)
// 	if action == "+" {
// 		fmt.Println(add(a, b))
// 	} else if action == "-" {
// 		fmt.Println(subst(a, b))
// 	} else if action == "*" {
// 		fmt.Println(mult(a, b))
// 	} else if action == "/" {
// 		fmt.Println(divid(a, b))
// 	}
// }

// func main() {
// 	var number int
// 	fmt.Println("Enter a number")
// 	fmt.Scan(&number)
// 	fmt.Println("Your even numbers: ")
// 	for i := 1; i <= number; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func main () {
// 	var number int
// 	fmt.Println("Enter a number to explore the fibonnaci's chain")
// 	fmt.Scanln(&number)
// 	a, b := 0, 1
// 	for i := 0; i <= number; i ++ {
// 		fmt.Print(a, " ")
// 		next := a + b
// 		a = b
// 		b = next
// 	}
// }

// func fibonacci(n int) []int {
// 	fib := make([]int, n)
// 	if n > 0 {
// 		fib[0] = 0
// 	}
// 	if n > 1 {
// 		fib[1] = 1
// 	}

// 	for i := 2; i < n; i++ {
// 		fib[i] = fib[i-1] + fib[i-2]
// 	}
// 	return fib
// }

// func main () {
// 	var n int
// 	fmt.Print("Enter a limit number to see a fibonacci chain: ")
// 	fmt.Scan(&n)
// 	result := fibonacci(n)
// 	for _, num := range result {
// 		fmt.Print(num, " ")
// 	}
// }

func main() {
	var n int
	fmt.Print("Enter a limit number to see prime numbers: ")
	fmt.Scanln(&n)
	fmt.Println("Here are the prime numbers up to", n, ": ")
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
