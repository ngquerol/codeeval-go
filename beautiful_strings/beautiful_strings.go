package main

import (
	"bufio"
	"unicode"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// Given a string, count how many times each of its characters (in the
// ASCII range [a-zA-Z]) appear. Returns a map containing these
// characters as keys, and their occurence count as values.
// Upper and lower case are considered the same.
func countCharOccurences(s string) map[rune]int {
	charOccurence := make(map[rune]int)

	for _, char := range s {

		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			lower_rune := unicode.ToLower(char)

			// if the character is already in the map, increment its
			// occurence count by 1; else, create a new entry for it,
			// with the occurence count set to 1.
			if _, present := charOccurence[lower_rune]; present {
				charOccurence[lower_rune] += 1
			} else {
				charOccurence[lower_rune] = 1
			}
		}
	}

	return charOccurence
}

// Given an array containing the occurence of the characters in a
// string, sorted in descending order, calculate the string's
// "beauty"; This consists in a integer between 1 and 26 (included).
// The most frequent characters are given the highest beauty, and each
// beauty is unique.
func calculateBeauty(occurences []int) int {
	total_beauty := 0

	// The most frequent character is given beauty 26, the second most
	// frequent one 25, and so on. Multiply each beauty by the number
	// of times the corresponding character appears in the string.
	for i, occurence := range occurences {
		total_beauty += (26 - i) * occurence
	}

	return total_beauty
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

		// count how many times each [a-zA-Z] character occurs in the
		// line
		charsOccurences := countCharOccurences(line)

		// keep only the occurences
		var occurences []int
		for _, occurence := range charsOccurences {
			occurences = append(occurences, occurence)
		}

		// sort them by descending order
		sort.Sort(sort.Reverse(sort.IntSlice(occurences)))

		// calculate the line's beauty score, based on the previously
		// found occurences of its characters
		fmt.Println(calculateBeauty(occurences))
	}
}
