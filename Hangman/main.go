package Hangman

import (
	a "HANGMAN_WEB_BURLE_CRAYSTON_LIERE/Hangman/fonctions"
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangManPositions []string
	UsedLetters      []string
}

func (h *HangManData) Init() {
	h.ToFind = a.RandomWord()
	h.Word = a.RevealLetter(h.ToFind)
	h.Attempts = 10
	h.HangManPositions = a.HangmanTab()
}

func (h *HangManData) Display() {
	fmt.Println("Mot à trouver:", h.Word)
	fmt.Println("Nombre d'essais restants:", h.Attempts)
}

func (h *HangManData) DisplayHangMan() {
	fmt.Println(h.HangManPositions[9-h.Attempts])
}

func (h *HangManData) Scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	for _, letter := range input {
		if len(input) == 1 && !unicode.IsLetter(letter) {
			h.Display()
			fmt.Println("Veuillez entrer une lettre !")
			h.Scan()
		}
	}
	if len(input) > 1 && len(input) != len(h.ToFind) {
		h.Display()
		fmt.Println("Veuillez entrer une seule lettre !")
		h.Scan()
	}
	if input == h.ToFind {
		h.Word = h.ToFind
		fmt.Println("Bravo, vous avez trouvé le mot:", h.Word)
	} else if len(input) == len(h.ToFind) {
		h.Attempts -= 2
		fmt.Println("Ce n'est pas le mot !")
		h.DisplayHangMan()
		h.Display()
		h.Scan()
	}
	if input == "" {
		h.Display()
		fmt.Println("Cette touche n'est pas disponible veuillez entrer une lettre !")
		h.Scan()
	}
	return input
}

func Hangman() {
	var hangman HangManData
	hangman.Init()
	for hangman.Attempts > 0 {
		hangman.Display()
		fmt.Println("Entrez une lettre ou un mot (-2 points si échec):")
		input := hangman.Scan()
		if a.VerifyLetter(input, hangman.ToFind) && len(input) == 1 {
			fmt.Println("La lettre", input, "n'est pas dans le mot !")
			hangman.Attempts--
			hangman.DisplayHangMan()
		}
		var indexes []int
		for index, letter := range hangman.ToFind {
			if input == string(letter) {
				if a.VerifyIndex(indexes, index) {
					indexes = append(indexes, index)
				}
			}
		}
		for _, index := range indexes {
			hangman.Word = a.Replace(hangman.Word, input, index)
		}
		if hangman.Word == hangman.ToFind {
			fmt.Println("Bravo, vous avez trouvé le mot:", hangman.Word)
			break
		}
	}
	if hangman.Attempts == 0 {
		fmt.Println("Vous avez perdu, le mot était:", hangman.ToFind)
	}
}
