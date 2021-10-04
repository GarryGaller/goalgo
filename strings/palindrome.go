package strings

import (
	"strings"
	"unicode"
)

func IsLetter(s string) bool {

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func IsPalindrome(s string) bool {

	// make the input string lower case.
	s = strings.ToLower(s)

	// make letters runes in case of multibyte not ASCII characters
	runes := []rune(s)
	start := 0
	end := len(runes) - 1

	for start < end {
		// if the letter at the start and end pointers are non-alphanumeric
		// then skip them.
		for !unicode.IsLetter(runes[start]) {
			start++
		}
		for !unicode.IsLetter(runes[end]) {
			end--
		}

		// return false immediately if two corresponding characters are not
		// matching.
		if runes[start] != runes[end] {
			return false
		}

		start++
		end--
	}

	return true
}
