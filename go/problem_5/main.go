package main

import "fmt"

func main() {
	for x := 2520; x < 1000000000; x += 2520 {
		if testNum(x) {
			fmt.Println("Problem 5:", x)
			break
		}
	}
}

func testNum(test int) bool {
	testSet := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for x := 0; x < 10; x++ {
		if test%testSet[x] != 0 {
			return false
		}
	}
	return true
}
