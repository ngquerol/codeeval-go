package main

import (
	"bufio"
	"fmt"
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

	var morse_alphabet = map[string]byte {
		".-":		'A',
		"-...":		'B',
		"-.-.":		'C',
		"-..":		'D',
		"."	:		'E',
		"..-.":		'F',
		"--.":		'G',
		"....":		'H',
		"..":		'I',
		".---":		'J',
		"-.-":		'K',
		".-..":		'L',
		"--":		'M',
		"-.":		'N',
		"---":		'O',
		".--.":		'P',
		"--.-":		'Q',
		".-.":		'R',
		"...":		'S',
		"-"	:		'T',
		"..-":		'U',
		"...-":		'V',
		".--":		'W',
		"-..-":		'X',
		"-.--":		'Y',
		"--..":		'Z',
		"-----":    '0',
		".----":    '1',
		"..---":    '2',
		"...--":    '3',
		"....-":    '4',
		".....":    '5',
		"-....":    '6',
		"--...":    '7',
		"---..":    '8',
		"----.":    '9',
	}

	// read the file line by line
	for scanner.Scan() {

		// get the string representation of the line
		line := scanner.Text()

		words := strings.Split(line, "  ")

		for word := range words {
			current_word := words[word];
			word_letters := strings.Split(current_word, " ")
			var translated_word []byte

			for _, letter := range word_letters {
				translated_word = append(translated_word, morse_alphabet[letter])
			}
			fmt.Printf("%s ", translated_word)
		}
		fmt.Println()
	}
}
