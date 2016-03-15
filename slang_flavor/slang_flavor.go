
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"io/ioutil"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("\nUsage: %s <text_file>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	text, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	s := string(text)

	var slang = [...]string {
		", yeah!",
		", this is crazy, I tell ya.",
		", can U believe this?",
		", eh?",
		", aw yea.",
		", yo.",
		"? No way!",
		". Awesome!",
	}

	odd := false
	replacements := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '?' || s[i] == '!' || s[i] == '.' {
			if odd {
				s = fmt.Sprintf("%s%s%s", s[:i], slang[replacements], s[i+1:])
				i = i + len(slang[replacements])
				replacements = (replacements + 1) % 8
				odd = false
			} else {
				odd = true
			}
		}
	}
	fmt.Println(s)
}
