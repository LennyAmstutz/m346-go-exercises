package main

import (
	"fmt"
	"time"
)

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month time.Month
	Year  int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName:         FullName{"Lenny", "Amstutz"},
		BirthDate:        BirthDate{20, 10, 2007},
		NumberOfSiblings: 1,
		ZodiacSign:       'W',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
