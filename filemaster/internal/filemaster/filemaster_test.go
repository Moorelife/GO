package filemaster

import (
	"testing"
)

var palindromes = []string{
	"negen",
	"kayak",
	"parterretrap",
	"",
}

func TestIsPalindromeTrue(t *testing.T) {
	for _, word := range palindromes {
		if !IsPalindrome(word) {
			t.Errorf("%s should have IsPalindrome() returning true", word)
		}
	}
}

var nonPalindromes = []string{
	"difficult",
	"dangerous",
	"edge case",
	"dummy",
}

func TestIsPalindromeFalse(t *testing.T) {
	for _, word := range nonPalindromes {
		if IsPalindrome(word) {
			t.Errorf("%s should have IsPalindrome() returning false", word)
		}
	}
}
