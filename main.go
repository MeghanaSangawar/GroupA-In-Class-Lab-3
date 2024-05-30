package main

// FahrenheitToCelsius converts a temperature from Fahrenheit to Celsius.
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Function for Converting temperature from Celsius to Farenheit
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}
