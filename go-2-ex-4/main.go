package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Students []Student
}

func main() {

	student1 := Student{FirstName: "Lenny", LastName: "Amstutz"}
	student2 := Student{FirstName: "Dylan", LastName: "Rodrigez"}
	student3 := Student{FirstName: "Lea", LastName: "Schuerch"}
	student4 := Student{FirstName: "Remo", LastName: "Albiser"}
	student5 := Student{FirstName: "Danial", LastName: "Najafi"}
	student6 := Student{FirstName: "Saskia", LastName: "Tellenbach"}

	class1 := Class{Students: []Student{student1, student2, student3}}
	class2 := Class{Students: []Student{student4, student5, student6}}

	module := map[string][]Class{
		"346": {class1},
		"111": {class1, class2},
		"222": {class2},
	}

	for moduleID, classes := range module {
		fmt.Printf("Modul %s wird von diesen Klassen besucht:\n", moduleID)
		for i, class := range classes {
			fmt.Printf("  Klasse %d:\n", i+1)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.FirstName, student.LastName)
			}
		}
		fmt.Println()
	}
}
