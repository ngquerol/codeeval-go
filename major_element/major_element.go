package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getBoyerMooreCandidate(elements []string) string {
	candidate := ""
	count := 0

	for i := range elements {
		if count == 0 {
			candidate = elements[i]
		}
		if candidate == elements[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func checkBoyerMooreCandidate(element string, elements []string) bool {
	majorityCount := (len(elements) / 2) + 1
	isMajor := false
	count := 0

	for i := range elements {
		if elements[i] == element {
			count++
		}
		if count >= majorityCount {
			isMajor = true
			break
		}
	}
	return isMajor
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

		elements := strings.Split(line, ",")

		majorityCandidate := getBoyerMooreCandidate(elements)

		if checkBoyerMooreCandidate(majorityCandidate, elements) {
			fmt.Println(majorityCandidate)
		} else {
			fmt.Println("None")
		}
	}
}
