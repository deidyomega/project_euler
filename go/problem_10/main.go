package main

import "fmt"

// const arrSize int = 10
const arrSize int = 2000000

func removal(arr *[arrSize]int, num int) {
	if num != 0 {
		for x := num + num; x < len(arr); x += num {
			arr[x] = 0
		}
	}
}

func main() {
	var primes [arrSize]int
	var total int
	for x := 2; x < arrSize; x++ {
		primes[x] = x
	}

	for x := 2; x < arrSize; x++ {
		removal(&primes, primes[x])
		total += primes[x]
	}

	fmt.Println("Problem 10:", total)
}
