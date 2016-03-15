package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

		list := strings.Split(scanner.Text(), ",")

		x, _ := strconv.Atoi(list[0])
		n, _ := strconv.Atoi(list[1])

		multiple := n

		for {
			if multiple > x {
				break
			}
			multiple += n
		}

		fmt.Println(multiple)
	}
}
