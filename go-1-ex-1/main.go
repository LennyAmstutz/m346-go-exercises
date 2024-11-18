package main

import "fmt"

func main() {

	var firstName string = "Lenny"
	var lastName string = "Lenny"

	var dayOfBirth int = 20
	var monthOfBirth int = 10
	var yearOfBirth int = 2007

	var numberOfSiblings int = 1

	var heightInMeters float32 = 1.74

	var zodiacSign string = "Wage"

	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
