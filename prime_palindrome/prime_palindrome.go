package main

import (
	"math"
	"fmt"
	"strconv"
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

func isPalindrome(n int) bool {
	n_str := strconv.Itoa(n)

	if len(n_str) % 2 == 0 {
		return false
	}
	for i := 0; i < len(n_str) / 2; i++ {
		if n_str[i] != n_str[(len(n_str) - i) - 1] {
			return false
		}
	}
	return true
}

func main() {
	var largestPalindrome int

	for i := 0; i <= 1000; i++ {
		if isPrime(i) && isPalindrome(i) {
			if i > largestPalindrome {
				largestPalindrome = i
			}
		}
	}
	fmt.Println(largestPalindrome)
}
