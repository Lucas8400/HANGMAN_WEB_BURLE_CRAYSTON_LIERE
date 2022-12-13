package fonctions

func RevealLetter(word string) string {
	index := len(word)/2 - 1
	var letter1 string
	var new_word string
	for i, letter := range word {
		if i == index {
			letter1 += string(letter)
		}
	}
	for _, letter := range word {
		if string(letter) == letter1 {
			new_word += string(letter)
		} else {
			new_word += "_"
		}

	}
	return new_word
}
