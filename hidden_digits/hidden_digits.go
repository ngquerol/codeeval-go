
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

		var hidden []rune
		for _, char := range line {
			if char >= 'a' && char <= 'j' {
				hidden = append(hidden, char - '1')
			} else if char >= '0' && char <= '9' {
				hidden = append(hidden, char)
			}
		}
		if len(hidden) > 0 {
			for _, digit := range hidden {
				fmt.Printf("%c", digit)
			}
		} else {
			fmt.Print("NONE")
		}
		fmt.Println()
	}
}
