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

		s := line

		if len(s) > 55 {

			s = s[:40]

			for i := len(s) - 1; i > 0; i-- {
				if s[i] == ' ' {
					s = s[:i]
					break
				}
			}
			s = s + "... <Read More>"
		}
		fmt.Println(s)
	}
}
