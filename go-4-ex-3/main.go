package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := math.Pow(b, 2) - 4*a*c
	if D > 0 {
		return []float64{
			(-b + math.Sqrt(D)) / (2 * a),
			(-b - math.Sqrt(D)) / (2 * a),
		}
	} else if D == 0 {
		return []float64{-b / (2 * a)}
	}
	return []float64{}
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1))
	fmt.Println(computeQuadraticFormula(2, 4, 2))
	fmt.Println(computeQuadraticFormula(3, 4, 2))
}
