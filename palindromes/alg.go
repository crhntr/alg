package palindromes

// Alg returns whether a given string word
// is a palindromes.
func Alg(word string) bool {
	n := len(word) - 1
	for i := 0; len(word) > 0 && i <= len(word)/2; i++ {
		if word[i] != word[n-i] {
			return false
		}
	}
	return true
}
