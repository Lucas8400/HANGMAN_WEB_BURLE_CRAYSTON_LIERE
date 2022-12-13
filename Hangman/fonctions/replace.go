package fonctions

func Replace(word string, input string, index int) string {
	var new_word string
	for i, letter := range word {
		if i == index {
			new_word += input
		} else {
			new_word += string(letter)
		}
	}
	return new_word
}
