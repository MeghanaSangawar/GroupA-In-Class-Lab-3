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

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt the user to enter a temperature or type 'exit' to quit

		fmt.Print("Enter temperature (e.g., 32 F or 100 C) or type 'exit' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		parts := strings.Fields(input)
		// If the input does not consist of two parts (temperature and unit), prompt the user for valid input and continue to the next iteration

		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter a temperature followed by a unit (e.g., 32 F or 100 C).")
			continue
		}

		tempStr, unit := parts[0], strings.ToUpper(parts[1])
		// Parse temperature string into a float64 value

		temp, err := strconv.ParseFloat(tempStr, 64)

		// If parsing fails (non-numeric input), prompt the user for valid input and continue to the next iteration

		if err != nil {
			fmt.Println("Invalid temperature. Please enter a numeric value.")
			continue
		}
		// Perform temperature conversion based on the unit provided by the user

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
