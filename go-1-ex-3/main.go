package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	eyesFile, _ := os.Create("eyes.txt")
	defer eyesFile.Close()

	diceFile, _ := os.Create("eyes.txt")
	defer diceFile.Close()

	fmt.Fprintln(eyesFile, "The dice shows", eyes, "eyes")
	fmt.Fprintln(diceFile, "The dice was rolled at", when)

	fmt.Println("the dice shows", eyes, "eyes")

	fmt.Println("the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO
}
