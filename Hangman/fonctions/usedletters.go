package fonctions

func UsedLetters(tab []string, input string) bool {
	for _, letter := range tab {
		if letter == input {
			return true
		}
	}
	return false
}
