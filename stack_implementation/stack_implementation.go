package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"errors"
	"strconv"
	"strings"
)

type intStack struct {
	Items [4096]int
	Size int
}

func (s *intStack) Push(i int) error {
	if s.Size == len(s.Items) {
		return errors.New("Stack is full, cannot push.")
	} else {
		s.Items[s.Size] = i
		s.Size++
		return nil
	}
}

func (s *intStack) Pop() (int, error) {
	if s.Size == 0 {
		return 0, errors.New("Stack is empty, cannot pop.")
	} else {
		ret := s.Items[s.Size-1]
		s.Items[s.Size-1] = 0
		s.Size--
		return ret, nil
	}
}

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

		s := intStack{}

		ints_str := strings.Split(line, " ")

		for _, int_str := range ints_str {
			int, _ := strconv.Atoi(int_str)
			s.Push(int)
		}

		i := 0

		for {
			if x, err := s.Pop(); err != nil {
				break
			} else if i % 2 == 0 {
				fmt.Print(x, " ")
			}
			i++
		}
		fmt.Println()
	}
}
