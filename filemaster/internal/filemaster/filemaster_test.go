package filemaster

import (
	"testing"
)

palindromes := []string{
	"negen"
	"kayak"
	"parterretrap"
	""
}

func TestIsPalindromeTrue(t *testing.T) {
	for word := range palindromes {
		if !IsPalindrome(word) {
			t.Error("%v should have IsPalindrome() returning true", word)
		}
	}
}

nonPalindromes := []string {
	"difficult"
	"dangerous"
	"edge case"
	"dummy"
}

func TestIsPalindromeFalse(t *testing.T) {
	for word := range palindromes {
		if IsPalindrome(word) {
			t.Error("%v should have IsPalindrome() returning true", word)
		}
	}
}
