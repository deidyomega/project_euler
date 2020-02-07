package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		stop bool
		a    float64
		b    float64
		c    float64
	)
	for c = 100; c < 1000; c++ {
		for b = 100; b < c; b++ {
			for a = 100; a < b; a++ {
				if a+b+c == 1000 {
					if math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2) {
						fmt.Printf("A: %.0f B: %.0f C: %.0f\n", a, b, c)
						fmt.Printf("Problem 6: %.0f\n", a*b*c)
						stop = true
					}
				}
				if stop {
					break
				}
			}
			if stop {
				break
			}
		}
		if stop {
			break
		}
	}
}
