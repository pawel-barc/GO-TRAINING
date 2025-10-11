package main

import (
	"fmt"
)

func main() {
	var height float32
	var weight float32
	var bmi float32
	fmt.Println("====BMI Calculator====")
	for {
		fmt.Println("Enter your height (m): ")
		fmt.Scanln(&height)
		if height <= 0 {
			fmt.Println("Your height can't be lower or equal 0")
		} else {
			break
		}
	}

	for {
		fmt.Println("Enter your weight(kg): ")
		fmt.Scan(&weight)
		if weight <= 0 {
			fmt.Println("Your weight can't be lower or equal 0")
		} else {
			break
		}
	}

	bmi = weight / (height * height)
	fmt.Println("Your BMI is ", bmi)

	if bmi < 18.5 {
		fmt.Println("You have underweight")
	} else if bmi < 25 {
		fmt.Println("You have normal weight")
	} else if bmi < 30 {
		fmt.Println("You have overweight")
	} else {
		fmt.Println("You are obese")
	}

}
