
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

		// get the number and the two positions from the string
		numbers := strings.Split(line, ",")

		number, _	:= strconv.Atoi(numbers[0])
		p1, _		:= strconv.Atoi(numbers[1])
		p2, _		:= strconv.Atoi(numbers[2])

		// get a binary representation of our number
		binary_number := fmt.Sprintf("%b", number)

		// compare bits at positions p1 & p2
		if binary_number[len(binary_number) - p1] ==
			binary_number[len(binary_number) - p2] {
			fmt.Println("true ")
		} else {
			fmt.Println("false")
		}
	}
}
