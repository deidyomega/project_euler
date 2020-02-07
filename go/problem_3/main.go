package main

import (
	"fmt"
	"sort"
)

const arrSize int = 10000

func removal(arr *[arrSize]int, num int) {
	if num != 0 {
		for x := num + num; x < len(arr); x += num {
			arr[x] = 0
		}
	}
}

func sum(arr *[arrSize]int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func prime_list() []int {
	var primes [arrSize]int
	for x := 2; x < arrSize; x++ {
		primes[x] = x
	}

	for x := 2; x < arrSize; x++ {
		removal(&primes, primes[x])
	}

	final := make([]int, 0)
	for x := 0; x < arrSize; x++ {
		if primes[x] != 0 {
			final = append(final, x)
		}
	}
	return final

}

func primeFactorization(num int, primes []int, pList *[]int) []int {
	for _, prime := range primes {
		if num%prime == 0 {
			*pList = append(*pList, prime)
			*pList = primeFactorization(num/prime, primes, pList)
			break
		}
	}
	return *pList
}

func main() {
	x := make([]int, 0)
	lst := primeFactorization(600851475143, prime_list(), &x)
	sort.Ints(lst) // Sorted
	fmt.Println("Problem 10:", lst[len(lst)-1])
}
