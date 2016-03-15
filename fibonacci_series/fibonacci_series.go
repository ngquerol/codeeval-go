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

		x, _ := strconv.ParseInt(line, 10, 8)

		fmt.Println(nthFibo(x))
	}
}

func nthFibo(x int64) int64 {

	if x == 0 {
		return x
	} else if x == 1 {
		return x
	} else {
		return nthFibo(x - 1) + nthFibo(x - 2)
	}
}
