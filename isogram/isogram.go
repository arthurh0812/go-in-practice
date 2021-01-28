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

// IsAnagram reports whether the given string is an anagram (meaning that it is "mirrowable")
func IsAnagram(s string) bool {
	reversed := make([]byte, len(s))
	for i, j := 0, len(s)-1; j >= 0; i, j = i+1, j-1 {
		reversed[i] = s[j]
	}
	return s == string(reversed)
}
