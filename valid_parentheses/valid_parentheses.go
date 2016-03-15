package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"errors"
)

type byteStack struct {
	Items [1024]byte
	Size int
}

func (s *byteStack) Push(b byte) error {
	if s.Size == len(s.Items) {
		return errors.New("Stack is full, cannot push.")
	} else {
		s.Items[s.Size] = b
		s.Size++
		return nil
	}
}

func (s *byteStack) Pop() (byte, error) {
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


	closings := map[byte]byte {
		'}': '{',
		']': '[',
		')': '(',
	}

	 openings := map[byte]byte {
		'{': '}',
		'[': ']',
		'(': ')',
	}

	scanner := bufio.NewScanner(file)

	// read the file line by line
	for scanner.Scan() {

		// get the string representation of the line
		line := scanner.Text()

		s := byteStack{}

		isValid := true

		for i := range line {
			if _, present := openings[line[i]]; present {
				s.Push(line[i])
			} else {
				x, _ := s.Pop()
				// if the closing char does not correspond to the last opened
				// char, the string is invalid.
				if x != closings[line[i]] {
					isValid = false
					break
				}
			}
		}
		// if the stack is not empty, there is unclosed chars left; the string
		// is thus invalid.
		if (isValid && s.Size == 0) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
