package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		square float64
		sum    float64
		x      float64
	)

	for x = 0; x <= 100; x++ {
		sum += x
		square += math.Pow(x, 2)
	}

	// %.0f forces to get rid of scientific notation
	fmt.Printf("Problem 6: %.0f\n", math.Pow(sum, 2)-square)
}
