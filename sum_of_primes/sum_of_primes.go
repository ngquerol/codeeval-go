package main

import (
	"math"
	"fmt"
)

func isPrime(n int) bool {
	if n <= 3 {
		return n >= 2
	}
	if n % 2 == 0 || n % 3 == 0 {
		return false
	}
	// TODO: test only if n is divisible by a prime numbers lesser
	// than n
	for i := 5; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 || n % (i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	i := 0
	primeCount := 0
	sumOfPrimes := 0

	for primeCount < 1000 {
		if isPrime(i) {
			sumOfPrimes += i
			primeCount++
		}
		i++
	}
	fmt.Println(sumOfPrimes)
}
