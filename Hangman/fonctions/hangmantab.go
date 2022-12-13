package fonctions

import (
	"bufio"
	"log"
	"os"
)

func HangmanTab() []string {
	content, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)
	var element string
	var lines []string
	for scanner.Scan() {
		if scanner.Text() != "" {
			element += "\n" + scanner.Text()
		} else {
			lines = append(lines, element)
			element = ""
		}
	}
	return lines
}
