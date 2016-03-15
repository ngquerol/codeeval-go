
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"bufio"
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

		sets := strings.Split(line, ";")

		set1 := strings.Split(sets[0], ",")
		set2 := strings.Split(sets[1], ",")

		var intersect []string

		for _, e1 := range set1 {
			for _, e2 := range set2 {
				if e1 == e2 {
					intersect = append(intersect, e1)
				}
			}
		}

		if len(intersect) != 0 {
			for i := range intersect {
				if i == len(intersect) - 1 {
					fmt.Print(intersect[i])
				} else {
					fmt.Print(intersect[i], ",")
				}
			}
		}
		fmt.Println()
	}
}
