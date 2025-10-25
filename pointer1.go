package main

import "fmt"

// func main() {
// 	x := 10
// 	p := &x
// 	fmt.Println("Address x:", p)
// 	fmt.Println("'p' value:", *p)

// 	*p = 20
// 	fmt.Println("'x' new value:", x)
// }

// func doubleValue(n *int) {
// 	*n = *n * 2
// }

// func main() {
// 	num := 5
// 	fmt.Println("Before:", num)
// 	doubleValue(&num)
// 	fmt.Println("After:", num)
// }

// func doubleValue(n *int) {
// 	*n = *n * 2
// }

// func main() {
// 	var num int
// 	fmt.Println("Enter number to double it five times")
// 	fmt.Scanln(&num)
// 	for i := 1; i <= 5; i++ {
// 		doubleValue(&num)
// 	}
// 	fmt.Println("The result is: ", num)
// }

// func deposit(b *float64, amount float64) {
// 	amount = 50
// 	*b = *b + amount
// }

// func main() {
// 	var balance float64
// 	balance = 100
// 	deposit(&balance, 50)
// 	fmt.Println("The result is: ", balance)

// }

// func main() {
// 	x := 10
// 	p := &x
// 	pp := &p

// 	fmt.Println("value x:", x)
// 	fmt.Println("Address x:", p)
// 	fmt.Println("Address p:", pp)
// 	fmt.Println("*p:", *p)
// 	fmt.Println("**pp:", **pp)

// 	**pp = 45

// 	fmt.Println("value x after modification: ", x)
// }

type Book struct {
	Title string
	Pages int
}

func main() {
	b := Book{Title: "Dune", Pages: 500}
	pb := &b
	fmt.Println("Title:", pb.Title)
	fmt.Printf("address of b: %p\n", pb)

	pb.Pages = 411

	fmt.Println("Pages after modif: ", b.Pages)
}

