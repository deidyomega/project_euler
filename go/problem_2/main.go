package main

import "fmt"

func main() {
	a := 1
	b := 1
	c := 0
	total := 0

	for c < 4000000 {
		c = a + b
		a = b
		b = c
		if c%2 == 0 {
			total += c
		}
	}
	fmt.Println("Problem 2:", total)
}
