package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

		var repeated string

		for i := range line {
			for j := i; j < len(line); j++ {
				if line[j] != line[i] {
					repeated = line[i:j]
					break
				}
				if j == len(line) - 1 {
					repeated = line[i:j+1]
					break
				}
			}
			if len(repeated) > 1 {
				line = strings.Replace(line, repeated, repeated[0:1], 1)
			}
		}
		fmt.Println(line)
	}
}
