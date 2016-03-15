package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"strconv"
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

		elements := strings.Split(line, ",")

		var words, digits []string

		for _, element := range elements {
			// do we have an integer or a string ?
			if _, err := strconv.Atoi(element); err != nil {
				digits = append(digits, element)
			} else {
				words = append(words, element)
			}
		}
		if len(digits) != 0 {
			for i, digit := range digits {
				if i == len(digits) - 1 {
					fmt.Print(digit)
				} else {
					fmt.Print(digit, ",")
				}
			}
		}
		if len(words) != 0 && len(digits) != 0 {
			fmt.Print("|")
		}
		if len(words) != 0 {
			for i, word := range words {
				if i == len(words) - 1 {
					fmt.Print(word)
				} else {
					fmt.Print(word, ",")
				}
			}
		}
		fmt.Println()
	}
}
