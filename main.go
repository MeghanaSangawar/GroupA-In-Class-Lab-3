package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FahrenheitToCelsius converts a temperature from Fahrenheit to Celsius.
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Function for Converting temperature from Celsius to Farenheit
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}
//Function main
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter temperature (e.g., 32 F or 100 C) or type 'exit' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter a temperature followed by a unit (e.g., 32 F or 100 C).")
			continue
		}

		tempStr, unit := parts[0], strings.ToUpper(parts[1])
		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			fmt.Println("Invalid temperature. Please enter a numeric value.")
			continue
		}

		switch unit {
		case "F":
			celsius := FahrenheitToCelsius(temp)
			fmt.Printf("%.2f F is %.2f C\n", temp, celsius)
		case "C":
			fahrenheit := CelsiusToFahrenheit(temp)
			fmt.Printf("%.2f C is %.2f F\n", temp, fahrenheit)
		default:
			fmt.Println("Invalid unit. Please enter 'F' for Fahrenheit or 'C' for Celsius.")
		}
	}
}
