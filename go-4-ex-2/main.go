package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(sideA float64, sideB float64) float64 {
	c := math.Sqrt(math.Pow(sideA, 2) + math.Pow(sideB, 2))
	return c
}

func main() {
	fmt.Println(computeHypotenuse(30, 40))
	fmt.Println(computeHypotenuse(50, 60))
	fmt.Println(computeHypotenuse(20, 30))
	fmt.Println(computeHypotenuse(25, 35))
}
