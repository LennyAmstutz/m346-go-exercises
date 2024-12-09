package main

import "fmt"

func convertCelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	c1, c2, c3 := 0.0, 25.0, 100.0

	f1 := convertCelsiusToFahrenheit(c1)
	f2 := convertCelsiusToFahrenheit(c2)
	f3 := convertCelsiusToFahrenheit(c3)

	fmt.Println("Celsius to Fahrenheit:")
	fmt.Println(c1, "°C ->", f1, "°F")
	fmt.Println(c2, "°C ->", f2, "°F")
	fmt.Println(c3, "°C ->", f3, "°F")

	fmt.Println("Fahrenheit back to Celsius:")
	fmt.Println(f1, "°F ->", convertFahrenheitToCelsius(f1), "°C")
	fmt.Println(f2, "°F ->", convertFahrenheitToCelsius(f2), "°C")
	fmt.Println(f3, "°F ->", convertFahrenheitToCelsius(f3), "°C")
}
