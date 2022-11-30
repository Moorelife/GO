package filemaster

func IsPalindrome(word string) bool {
	for i := range word {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}
