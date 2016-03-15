package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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

		// extract all digits from the line
		var elements []int

		for _, digit := range line {
			elements = append(elements, int(digit) - 48)
		}

		retVal := 1

		for i := range elements {
			if countDigit(elements, i) != int(elements[i]) {
				retVal = 0
				break
			}
		}
		fmt.Println(retVal)
	}
}

func countDigit(elements []int, x int) int {
	count := 0

	for _, digit := range elements {
		if digit == x {
			count++
		}
	}
	return count
}
