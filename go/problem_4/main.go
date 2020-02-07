package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		answer int
		tmp    int
		x      int
		y      int
	)

	for x = 900; x < 1000; x++ {
		for y = 990; y < 1000; y++ {
			tmp = x * y
			stmp := strconv.Itoa(tmp)

			if reverse(stmp) == stmp {
				if tmp > answer {
					answer = tmp
				}
			}
		}
	}
	fmt.Println("Problem 1:", answer)
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
