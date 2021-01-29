package isogram

import "unicode"

// IsIsogram reports whether the given string is an isogram (meaning that each character is unique)
func IsIsogram(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		if !seen[r] {
			seen[r] = true
			continue
		}
		return false
	}
	return true
}

// IsPalindrome reports whether the given string is an anagram (meaning that it is "mirrowable")
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	last := len(letters) - 1
	for i := range letters {
		if letters[i] != letters[last-i] {
			return false
		}
	}
	return true
}
