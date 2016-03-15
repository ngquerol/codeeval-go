package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	scanner := bufio.NewScanner(file)

	// read the file line by line
	for scanner.Scan() {

		// get the string representation of the line
		line := scanner.Text()
		// get all numbers in the line
		numbers_str := strings.Split(line, " ")

		// build an array from the file's string
		var numbers []int

		for i := range numbers_str {
			var x int
			fmt.Sscanf(numbers_str[i], "%d", &x)
			numbers = append(numbers, x)
		}

		// for large sets of numbers, it would be better if we
		// progressively built an array containing all non-uniques
		// numbers, and check that one before checking the entire set.
		// then we would not need to do it for each number.
		lowest := 10
		lowest_idx := -1

		for i, number := range numbers {
			unique := true
			for j, next_number := range numbers {
				if i != j && number == next_number {
					unique = false
				}
			}
			if unique && number < lowest {
				lowest = number
				lowest_idx = i
			}
		}
		fmt.Println(lowest_idx + 1)
	}
}
