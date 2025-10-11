// package main

// import "fmt"

// func main() {
// 	var shoppingList [] string
// 	var choice string

// 	fmt.Println("===Welcome to your Shopping List===")

// 	for {
// 		fmt.Println("\nChoose an action\n")
// 		fmt.Println("1 - add an item \n")
// 		fmt.Println("2 - show the list\n")
// 		fmt.Println("3 - remove an item\n")
// 		fmt.Println("4 - exit\n")
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case "1":
// 			var item string
// 			fmt.Println("Enter an item: ")
// 			fmt.Scanln(&item)
// 			shoppingList = append(shoppingList, item)
// 			fmt.Printf("\n%s added to the list\n", item)

// 		case "2":
// 			if len(shoppingList) == 0 {
// 				fmt.Println("Your list is empty")
// 			} else {
// 				fmt.Println("Your shopping list: \n")
// 				for i, item := range shoppingList {
// 					fmt.Printf("%d. %s\n", i+1, item)
// 				}
// 			}

// 		case "3":
// 			if len(shoppingList) == 0 {
// 				fmt.Println("Nothing to remove, your list is empty")
// 				continue
// 			}
// 			fmt.Println("Enter a number of the item to remove: ")
// 			for i, item := range shoppingList {
// 				fmt.Printf("%d. %s\n", i+1, item)
// 			}
// 			var index int
// 			fmt.Scanln(&index)
// 			if index < 1 || index > len(shoppingList) {
// 				fmt.Println("Wrong number! Try again")
// 				continue
// 			}
// 			removed := shoppingList[index - 1]
// 			shoppingList = append(shoppingList[:index - 1], shoppingList[index:]...)
// 			fmt.Println(removed, "was removed.")
// 			break

// 		case "4":
// 			fmt.Println("Goodbye")
// 			return

// 		default:
// 			fmt.Println("Unknown choice! Try again")
// 		}
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var shoppingList []string
	var options int

	shoppingList = loadList("shopping_list.txt")

	fmt.Println("=== Welcome to your Shopping List ===")

	for {
		fmt.Println("Your options:\n ")
		fmt.Println("1 - Add an item")
		fmt.Println("2 - Show the list")
		fmt.Println("3 - Remove an item")
		fmt.Println("4 - Save and exit")
		fmt.Print("Enter one option: ")

		fmt.Scanln(&options)

		switch options {
		case 1:
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter an item")
			item, _ := reader.ReadString('\n')
			item = strings.TrimSpace(item)
			shoppingList = append(shoppingList, item)
			fmt.Printf("%s added to the list\n", item)

		case 2:
			if len(shoppingList) == 0 {
				fmt.Println("Your list ist empty")
			} else {
				fmt.Println("Your shopping list: ")
				for i, item := range shoppingList {
					fmt.Printf("%d. %s\n", i+1, item)
				}
			}

		case 3:
			var index int
			fmt.Println("Remove an item using identification number: ")
			if len(shoppingList) == 0 {
				fmt.Println("Your list is empty")
				continue
			}
			for i, item := range shoppingList {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			fmt.Scanln(&index)
			if index < 1 || index > len(shoppingList) {
				fmt.Println("Unknown item, try again ")
				continue
			}
			removed := shoppingList[index-1]
			shoppingList = append(shoppingList[:index-1], shoppingList[index:]...)
			fmt.Printf("%s removed\n", removed)

		case 4:
			saveList("shopping_list.txt", shoppingList)
			fmt.Println("List saved. Goodbye!!")
			return

		default:
			fmt.Println("Unknown choice, try again.")

		}

	}
}

func loadList(filename string) []string {
	var list []string
	
	file, err := os.Open(filename)
	if err != nil {
		return list
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}

func saveList(filename string, list []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("An error occured by saving", err)
		return 
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, item := range list {
		writer.WriteString(item + "\n")
	}
	writer.Flush()
}