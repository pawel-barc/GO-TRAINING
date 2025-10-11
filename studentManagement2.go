package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name string
	Age int
	Average float64
}

func main() {
	var students []Student
	var option int
	reader := bufio.NewReader(os.Stdin)
	students = loadStudents("students.txt")

	fmt.Println("=== Welcome to Students Management Program ===")
	for {
		fmt.Println("Using numbers select an option from below :")
		fmt.Println("1 - Add student")
		fmt.Println("2 - Show all students")
		fmt.Println("3 - Show grade average of a student")
		fmt.Println("4 - Save end exit")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var s Student
			
			fmt.Println("Name:")
			name, _ := reader.ReadString('\n')
			s.Name = strings.TrimSpace(name)

			fmt.Println("Age:")
			fmt.Scanln(&s.Age)

			fmt.Println("Average:")
			fmt.Scanln(&s.Average)

			students = append(students, s)
			fmt.Printf("%s added!\n", s.Name)

		case 2:
			if len(students) == 0 {
				fmt.Println("No students yet")
				continue
			}
			fmt.Println("List of students: ")
			for i, s := range students {
				fmt.Printf("%d. %s: (Age: %d) - Average: %.2f\n", i+1, s.Name, s.Age, s.Average)
			}
		
		case 3:
			fmt.Println("Choose a student from the list: ")
			for i, s := range students {
				fmt.Printf("%d.  %s\n", i+1, s.Name)
				if len(students) == 0 {
					fmt.Println("No students yet")
					continue
				}
			}
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			found := false
			for _, s := range students {
				if s.Name == name {
					fmt.Printf("Average grade for %s is: %.2f\n", s.Name, s.Average)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Student not found")
			}

		case 4:
			saveStudents("students.txt", students)
			fmt.Println("Goodbye")
			return
			
		default:
			fmt.Println("Unknown option, try again")	
		} 
	}
}

func loadStudents(filename string) []Student {
	var list []Student

	file, err := os.Open(filename)
	if err != nil {
		println("No file found")
		return list
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			var s Student
			s.Name = parts[0]
			fmt.Sscanf(parts[1], "%d", &s.Age)
			fmt.Sscanf(parts[2], "%.2f", &s.Average)
			list = append(list, s)
		}
	}
	return list
}

func saveStudents(filename string, list []Student) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error by saving file: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, s := range list {
		fmt.Fprintf(writer, "%s,%d,%.2f\n", s.Name, s.Age, s.Average)
	}
	writer.Flush()
}