package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	students := make(map[string][]float64)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== Welcome to the student management program ===")
		fmt.Println("Chose one of below options")
		fmt.Println("1 - Add a student and grades")
		fmt.Println("2 - Show all students")
		fmt.Println("3 - Show an average for a student")
		fmt.Println("4 - Exit")
		fmt.Print("Choose option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		
		switch option {
		case "1":
			fmt.Println("Add a student name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("Enter grades separated by spaces: ")
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			parts := strings.Split(line, " ")

			var grades []float64
			for _, grade := range parts {
				num, err := strconv.ParseFloat(grade, 64)
				if err != nil {
					fmt.Println("Invalid grade!")
					continue
				}
				grades = append(grades, num)
			}
			students[name] = grades
			fmt.Printf("Student %s added with grades: %v\n", name, grades)
			
		case "2":
			fmt.Println("Students List: ")
			if len(students) == 0 {
				fmt.Println("No students yet")
				continue
			}
			for name, grades := range students {
				fmt.Println(name, ":", grades)
			}

		case "3":
			if len(students) == 0 {
				fmt.Println("Student list is empty!")
				continue
			}
			fmt.Println("Enter the name of a student")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			grades, exists := students[name]
			if !exists {
				fmt.Println("Student not found")
				continue
			}
			avg := calculateAverage(grades)
			fmt.Printf("Average grade for %s is: %.2f\n", name, avg) 			
		
		case "4":
			fmt.Println("See ya")
			return
			
		default:
			fmt.Println("Unknown option, try again!")	
		}
	}
}

func calculateAverage(grades []float64) float64 {
	if len(grades) == 0 {
		return 0
	}
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}
