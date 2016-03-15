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

		// get the string representation of the line
		line := scanner.Text()

		pairs := strings.Split(line, " ")
		var distances []int

		for _, pair := range pairs {
			city := strings.Split(pair, ",")
			distance_str := city[1][0:len(city[1]) - 1]
			distance, _ := strconv.Atoi(distance_str)
			distances = append(distances, distance)
		}

		for i := range distances {

			min := 30001

			for i := range distances {
				if distances[i] < min && distances[i] != 0 {
					min = distances[i]
				}
			}

			for i := range distances {
				if distances[i] != 0 {
					distances[i] -= min
				}
			}

			if i != (len(distances) - 1) {
				fmt.Print(min,",")
			} else {
				fmt.Print(min)
			}
		}
		fmt.Println()
	}
}
