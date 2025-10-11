package main

import (
	"fmt"
)

func main() {
	var kelvin float64
	var farenheit float64
	var celsius float64

	fmt.Println("Enter temperature in Celsius ")
	fmt.Scan(&celsius)
	kelvin = celsius + 273.15
	farenheit = celsius*9/5 + 32
	fmt.Printf("\n%.2f°C equals %.2f°F and %.2f K", celsius, farenheit, kelvin)
}
