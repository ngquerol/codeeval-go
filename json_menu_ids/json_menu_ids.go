package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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

	// this data structure is built from the data contained in the
	// JSON stream
	type Data struct {
		Menu struct {
			Header string `json:"header"`
			Items  []struct {
				ID    int    `json:"id"`
				Label string `json:"label"`
			} `json:"items"`
		} `json:"menu"`
	}

	// read the file line by line
	for scanner.Scan() {

		// get the string representation of the line
		line := scanner.Text()

		// check if there is anything to decode
		if len(line) > 0 {
			dec := json.NewDecoder(strings.NewReader(line))
			sum := 0
			for {
				// decode the JSON stream and store the data in our struct
				var d Data
				if err := dec.Decode(&d); err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}
				// for each item that has a label, add its ID to the sum
				for _, item := range d.Menu.Items {
					if item.Label != "" {
						sum += item.ID
					}
				}
			}
			fmt.Println(sum)
		}
	}
}
