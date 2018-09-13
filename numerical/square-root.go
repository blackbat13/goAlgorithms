package main

import (
	"fmt"
	"math"
)

// Compute square root of n with precision p using Heron method
func squareRoot(n float64, p float64) float64 {
	x1 := n / 2
	x2 := (x1 + n/x1) / 2
	for math.Abs(x2-x1) > p {
		x1 = (x2 + n/x2) / 2
		x1, x2 = x2, x1
	}

	return x2
}

func main() {
	fmt.Println(squareRoot(9, 0.00001))
}
