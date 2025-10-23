package main
import "fmt"

func deposit(balance *float64, amount float64) {
	fmt.Println("Enter amount")
	fmt.Scanln(&amount)
	*balance += amount
}

func withdraw(balance *float64, amount float64) {
	fmt.Println("Enter amount")
	fmt.Scanln(&amount)
	if *balance >= amount {
		*balance -= amount
		fmt.Println("Withdrawal successful! ")
	} else {
		fmt.Println("Withdrawal not possible, not enought founds on your account")
	}
}

func main() {
	var myBalance float64 = 100

	fmt.Println("Initial balance: ", myBalance)
	var option int
	fmt.Println("Choose an option using numbers '1' and '2'!")
	fmt.Println("For deposit choose '1'")
	fmt.Println("For withdrawal choose '2'")
	fmt.Scanln(&option)

	switch option {
	case 1:
		deposit(&myBalance, amount)
		fmt.Printf("After deposit: \t%.2f (address: %p)\n", myBalance, &myBalance)

	case 2:
		withdraw(&myBalance, amount)
		fmt.Println("After withdrawal:", myBalance)

	default:
		fmt.Println("Unknown option")
	}

}