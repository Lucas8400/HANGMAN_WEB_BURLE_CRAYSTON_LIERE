package fonctions

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func RandomWord() string {
	readFile, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	rand.Seed(time.Now().Unix())
	mot := lines[rand.Intn(len(lines))]
	return mot
}
