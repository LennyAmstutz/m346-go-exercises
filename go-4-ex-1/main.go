package main

import (
	"errors"
	"fmt"
)

func berechneNote(erreichtePunkte, maxPunkte float64) (float64, error) {
	if erreichtePunkte < 0 || maxPunkte <= 0 {
		return 0, errors.New("Punkte müssen positiv sein")
	}
	if erreichtePunkte > maxPunkte {
		return 0, errors.New("erreichte Punkte dürfen nicht größer als max. Punkte sein")
	}

	note := 1 + 5*(1-(erreichtePunkte/maxPunkte))

	return note, nil
}

func main() {

	note1, err1 := berechneNote(17.5, 28.0)
	if err1 != nil {
		fmt.Println("Fehler:", err1)
	} else {
		fmt.Println("Note:", note1)
	}

	note2, err2 := berechneNote(25.0, 25.0)
	if err2 != nil {
		fmt.Println("Fehler:", err2)
	} else {
		fmt.Println("Note:", note2)
	}

	note3, err3 := berechneNote(-5.0, 20.0)
	if err3 != nil {
		fmt.Println("Fehler:", err3)
	} else {
		fmt.Println("Note:", note3)
	}
}
