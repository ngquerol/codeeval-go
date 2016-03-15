
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"bufio"
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

		lists := strings.Split(line, " | ")

		list1 := strings.Split(lists[0], " ")
		list2 := strings.Split(lists[1], " ")

		for i := range list1 {
			n1, _ := strconv.Atoi(list1[i])
			n2, _ := strconv.Atoi(list2[i])
			fmt.Print(n1 * n2, " ")
		}
		fmt.Println()
	}
}
