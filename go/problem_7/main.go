package main

import (
	"fmt"
)

func main() {
	var seekNum = 10001
	gen := genPrime()
	for i := 1; i < seekNum; i++ {
		gen()
		if i == seekNum-1 {
			fmt.Println(gen())
		}
	}
	gen = nil // release for garbage collection
}

func genPrime() func() int {
	var (
		currentNumber = 2
		primes        []int
	)
	return func() int {
		var (
			isPrime bool
			prime   int
		)
		for true {
			isPrime = true
			for x := 0; x < len(primes); x++ {
				prime = primes[x]
				if prime*prime > currentNumber {
					break
				}
				if currentNumber%prime == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primes = append(primes, currentNumber)
				currentNumber++
				return currentNumber - 1
			}
			currentNumber++
		}
		return 0 // never hit
	}
}
