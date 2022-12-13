package fonctions

func VerifyIndex(tab []int, index int) bool {
	for _, element := range tab {
		if index == element {
			return false
		}
	}
	return true
}
