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

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read the file line by line
	for scanner.Scan() {

		// get the string representation of the line
		line := scanner.Text()

		// get all numbers in the line
		numbers_str := strings.Split(line, ",")

		// build an array from the file's string
		var numbers []int
		prev_x := -1 // set previous number to some non-existing value

		for _, number := range numbers_str {
			var x int

			fmt.Sscanf(number, "%d", &x)

			// if the number differs from the previous one, then it is
			// unique.
			if x != prev_x {
				numbers = append(numbers, x)
			}
			prev_x = x;
		}

		for i := range numbers {
			if i == len(numbers) - 1 {
				fmt.Printf("%d", numbers[i])
			} else {
				fmt.Printf("%d,", numbers[i])
			}
		}
		fmt.Println()
	}
}
