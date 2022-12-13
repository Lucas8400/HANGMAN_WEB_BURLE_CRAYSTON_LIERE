package fonctions

func VerifyLetter(input string, word string) bool {
	for _, letter := range word {
		if input == string(letter) {
			return false
		}
	}
	return true
}
