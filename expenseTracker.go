package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Expense struct {
	Description string
	Category string
	Amount float64
}

func(e Expense) Display() {
	fmt.Printf("Description: %s\nCategory: %s\nAmount: %.2f\n", e.Description, e.Category, e.Amount )
}

func main() {
	var expenses []Expense
	var option int
	reader := bufio.NewReader(os.Stdin)

	expenses = loadExpenses("expenses.txt")
	
	for {
		fmt.Println("==== Welcome to your app 'Expense Tracker' ====")
		fmt.Println("Please choose one option from below:")
		fmt.Println("1 - Add new expense")
		fmt.Println("2 - Show expenses list")
		fmt.Println("3 - Show calculated sum of all expenses")
		fmt.Println("4 - Show calculated sum of a category")
		fmt.Println("5 - Save and exit")

		fmt.Scanln(&option)

		switch option {
		case 1:
			var e Expense
			fmt.Println("Enter description")
			description, _ := reader.ReadString('\n')
			e.Description = strings.TrimSpace(description)

			fmt.Println("Enter category")
			category, _ := reader.ReadString('\n')
			e.Category = strings.TrimSpace(category)

			fmt.Println("Enter Prise ($)")
			var amount float64
			fmt.Scanln(&amount)
			e.Amount = amount 

			expenses = append(expenses, e)
			fmt.Println("Expense successfully added to the list!")
		
		case 2:
			if len(expenses) == 0 {
				fmt.Println("No expenses yet!")
				continue
			}
			fmt.Println("Here are your expenses:")
			for i, e := range expenses {
				fmt.Printf("%d.", i+1)
				e.Display()
			}

		case 3:
			var total float64
			
			if len(expenses) == 0 {
				fmt.Println("No expenses yet")
				continue
			}

			for _, e := range expenses {
				total += e.Amount
			}

			fmt.Printf("Total sum of all expenses: $%.2f\n", total)

		case 4:
			if len(expenses) == 0 {
				fmt.Println("No expenses yet")
				continue
			}
			fmt.Println("Enter category name to calculate total!")
			name, _ := reader.ReadString('\n')
			category := strings.TrimSpace(name)
			
			var total float64

			for _, e := range expenses {
				if strings.EqualFold(e.Category, category) {
					total += e.Amount
				}
			}

			if total == 0 {
				fmt.Println("No expenses for this category")
				continue
			} else {
				fmt.Printf("Total sum expenses for %s: $%.2f\n",category, total)
			}

		case 5:
			saveExpenses("expenses.txt", expenses)
			fmt.Println("Expenses saved! Goodbye!")
			return
		
		default:
			fmt.Println("Unknown option")	
		}

	}
}

func loadExpenses(filename string) []Expense {
	var list []Expense

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Imposible to open file", err)
		return list
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			amount, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				fmt.Println("Convertion error")
				continue
			}
			list = append(list, Expense{
				Description: parts[0],
				Category: parts[1],
				Amount: amount,
			})
		}

	}
	return list
}

func saveExpenses(filename string, list []Expense) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Impossible to save file", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, e := range list {
		fmt.Fprintf(writer, "%s,%s,%.2f\n", e.Description, e.Category, e.Amount)
	}
	writer.Flush()
	fmt.Printf("Data successfully saved to %s\n", filename) 
}