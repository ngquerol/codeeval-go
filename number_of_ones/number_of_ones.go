package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

		n, _ := strconv.Atoi(line)

		n_ones := 0

		if n == 0 {
			fmt.Print(0)
		} else {
			for n > 0 {
				if (n % 2 == 1) {
					n_ones++
				}
				n /= 2
			}
		}
		fmt.Println(n_ones)
	}
}
