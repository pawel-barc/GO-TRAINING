package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Title string
	Author string
	Pages int
	IsRead bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var library []Book
	var option int
	library = loadLibrary("library.txt")
	fmt.Println("=== Welcome to Book Manager Program ===")

	for {
		fmt.Println("Choose one of options:")
		fmt.Println("1 - Add a book")
		fmt.Println("2 - Show all books")
		fmt.Println("3 - Change read status")
		fmt.Println("4 - Save and goodbye ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var b Book
			fmt.Println("Enter book title:")
			title, _ := reader.ReadString('\n')
			b.Title = strings.TrimSpace(title)

			fmt.Println("Book author:")
			author, _ := reader.ReadString('\n')
			b.Author = strings.TrimSpace(author)

			fmt.Println("Total pages: ")
			fmt.Scanln(&b.Pages)

			b.IsRead = false

			library = append(library, b)
			fmt.Printf("Book %s successfully added:\n", b.Title)

		case 2:
			if len(library) == 0 {
				fmt.Println("No books yet")
				continue
			}
			fmt.Println("List of books:")
			for i, b := range library {
			status := "not read"
			if b.IsRead {
				status = "read"
			}
			fmt.Printf("%d. %s by %s (%d pages) - [%s]\n", i+1, b.Title, b.Author, b.Pages, status )
			}
		case 3:
			if len(library) == 0 {
				fmt.Print("No books in the library")
				continue
			}

			fmt.Println("Enter number of the book to mark as read")
			for i, b := range library {
				fmt.Printf("%d. %s [%t]\n", i+1, b.Title, b.IsRead)
			}
			var index int
			fmt.Scanln(&index)

			if index < 1 || index > len(library) {
				fmt.Println("Invalid number")
				continue
			}

			library[index-1].IsRead = true
			fmt.Printf("Book: '%s' is marked as read\n", library[index-1].Title)

		case 4:
			saveLibrary("library.txt", library)
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown option chosed, try again")
		}		
	}
}

func loadLibrary(filename string) []Book {
	var list []Book
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 4 {
			var b Book
			b.Title = parts[0]
			b.Author = parts[1]
			fmt.Sscanf(parts[2], "%d", &b.Pages)
			fmt.Sscanf(parts[3], "%t", &b.IsRead)
			list = append(list, b) 
		}
	}
	return list
}

func saveLibrary(filename string, list []Book) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("File couln't be save: ", err)
		return
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, b := range list {
		fmt.Fprintf("%s,%s,%d,%t\n", b.Title, b.Author, b.Pages, b.IsRead )
	}
	writer.Flush()
}