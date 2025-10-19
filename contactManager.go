package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Contact struct {
	Name string
	Phone string
	Email string
}

func (c Contact) Display() {
	fmt.Printf("Name: %s\n Phone number: %s\n Email: %s\n", c.Name, c.Phone, c.Email)
}

func isValidName(name string) bool {
	return len(name) >= 2
}

func isValidPhone(phone string) bool {
	clean := strings.ReplaceAll(phone, " ", "")
	clean = strings.ReplaceAll(phone, "-", "")

	match, _ := regexp.MatchString(`^\+?[0-9]{9,15}$`, clean)
	return match
}

func isValidEmail(email string) bool {
	pattern := `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func saveContacts(filename string, list []Contact) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error by file saving", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, c := range list {
		fmt.Fprintf(writer, "%s,%s,%s\n", c.Name, c.Phone, c.Email)
	}
	writer.Flush()
	fmt.Printf("Contacts successfully saved to %s\n", filename)
}

func loadContacts(filename string) []Contact {
	var list []Contact

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Impossible to open file", err)
		return list
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			list = append(list, Contact{
				Name: parts[0],
				Phone: parts[1],
				Email: parts[2],
			})
		}

	}
	return list

}

func main () {
	var contacts []Contact
	var option int
	reader := bufio.NewReader(os.Stdin)

	contacts = loadContacts("contacts.txt")

	for {
		fmt.Println("=== Welcome to Your 'Contacts Manager' ===")
		fmt.Println("Choose an option from below")
		fmt.Println("1 - Add contact")
		fmt.Println("2 - Show contacts list")
		fmt.Println("3 - Search contact by name")
		fmt.Println("4 - Remove contact")
		fmt.Println("5 - Save and exit")

		fmt.Scanln(&option)

		switch option {
		case 1:
			var c Contact
			fmt.Println("Enter name")
			name, _ := reader.ReadString('\n')
			c.Name = strings.TrimSpace(name)

			if !isValidName(c.Name) {
				fmt.Println("The name must contain at least 2 characteres")
				continue
			}

			fmt.Println("Enter phone number")
			phone, _ := reader.ReadString('\n')
			c.Phone = strings.TrimSpace(phone)

			if !isValidPhone(c.Phone) {
				fmt.Println("The phone number is not valid")
				continue
			}
			
			fmt.Println("Email")
			email, _ := reader.ReadString('\n')
			c.Email = strings.TrimSpace(email)

			if !isValidEmail(c.Email) {
				fmt.Println("The email is not valid")
				continue
			}
			contacts = append(contacts, c)
			fmt.Println("Contact successfully added!")

		case 2:
			if len(contacts) == 0 {
				fmt.Println("Your contacts list is empty")
				continue
			}
			fmt.Println("Contacts list:")
			for i, c := range contacts {
				fmt.Printf("%d. ", i+1)
				c.Display()
			}

		case 3:
			var search string
			fmt.Println("Enter name to search")
			name, _ := reader.ReadString('\n')
			search = strings.TrimSpace(name)
			found := false
			for _, c := range contacts {
				if strings.EqualFold(c.Name, search) {
					c.Display()
					found = true
				}
			}
			if !found {
				fmt.Println("Contact not found")
			}
			
		case 4:
			if len(contacts) == 0 {
				fmt.Println("No contacts to remove")
				continue
			}
			fmt.Println("Select a contact to remove:")
			for i, c := range contacts {
				fmt.Printf("%d. %s\n", i+1, c.Name)
			}
			fmt.Println("Enter a name")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			index := -1
			for i, c := range contacts {
				if strings.EqualFold(name, c.Name) {
					index = i
					break
				}
			}
			if index == -1 {
					fmt.Println("Contact not found")
				} else {
					removed := contacts[index].Name
					contacts = append(contacts[:index], contacts[index+1:]...)
					fmt.Printf("Contact: %s removed", removed)
				}	
			
		case 5:
			saveContacts("contacts.txt", contacts)
			fmt.Println("Goodbye!")
			return
		
		default:
			fmt.Println("Unknown option")	
		}
	}

}