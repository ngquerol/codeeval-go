package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"math"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 3 {
		return n >= 2
	}
	if n % 2 == 0 || n % 3 == 0 {
		return false
	}
	// TODO: test only if n is divisible by prime numbers lesser than
	// n
	for i := 5; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 || n % (i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("\nUsage: %s <text_file>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read the file line by line
	for scanner.Scan() {

		// get the string representation of the line
		line := scanner.Text()

		n, _ := strconv.Atoi(line)

		var prime_numbers []int
		i := 0

		for i < n {
			if isPrime(i) {
				prime_numbers = append(prime_numbers, i)
			}
			i++
		}

		for i, p := range prime_numbers {
			if i == (len(prime_numbers) - 1) {
				fmt.Print(p)
			} else {
				fmt.Print(p, ",")
			}
		}
		fmt.Println()
	}
}
