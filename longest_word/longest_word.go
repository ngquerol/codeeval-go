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

		words := strings.Split(line, " ")
		max_len := 0
		max_len_idx := -1

		for i, word := range words {
			word_len := len(word)

			if word_len > max_len {
				max_len = word_len
				max_len_idx = i
			}
		}
		fmt.Println(words[max_len_idx])
	}
}
