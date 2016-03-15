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

	const checkpoint = 'C'
	const gate = '_'
	last_idx := 0

	// read the file line by line
	for scanner.Scan() {

		// get the string representation of the line
		line := scanner.Text()

		foundCheckpoint := false

		for i := range line {
			if line[i] == checkpoint {
				line = line[:i] + getTurn(last_idx-i) + line[i+1:]
				last_idx = i
				foundCheckpoint = true
				break
			}
		}
		if !foundCheckpoint {
			for j := range line {
				if line[j] == gate {
					line = line[:j] + getTurn(last_idx-j) + line[j+1:]
					last_idx = j
					break
				}
			}
		}
		fmt.Println(line)
	}
}

func getTurn(i int) string {
	var turn string

	switch {
	case i == -1:
		turn = "\\"
	case i == 1:
		turn = "/"
	default:
		turn = "|"
	}
	return turn
}
