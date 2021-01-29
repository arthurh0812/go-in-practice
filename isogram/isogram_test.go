package isogram

import "testing"

func TestIsogram(t *testing.T) {
	// test input words that should be evaluated as an isogram and fail if IsIsogram() reports otherwise
	isograms := []string{"isogram", "palindrome"}
	for _, w := range isograms {
		if !IsIsogram(w) {
			t.Errorf(`IsIsogram(%q) = false`, w)
		}
	}
}

func TestNonIsogram(t *testing.T) {
	// test input words that shouldn't be evaluated as an isogram and fail if IsIsogram() reports otherwise
	nonIsograms := []string{"copper", "bicycle"}
	for _, w := range nonIsograms {
		if IsIsogram(w) {
			t.Errorf(`IsIsogram(%q) = true`, w)
		}
	}
}

func TestPalindrome(t *testing.T) {
	// test input words that should be evaluated as a palindrome and fail if IsPalindrome() reports otherwise
	palindromes := []string{"detartrated", "kayak"}
	for _, w := range palindromes {
		if !IsPalindrome(w) {
			t.Errorf(`IsPalindrome(%q) = false`, w)
		}
	}
}

func TestNonPalindrome(t *testing.T) {
	// test input words that shoudn't be a evaluated as a palidrome and fail if IsPalindrome() reports otherwise
	nonPalindromes := []string{"palindrome", "workbench"}
	for _, w := range nonPalindromes {
		if IsPalindrome(w) {
			t.Errorf(`IsPalindrome(%q) = true`, w)
		}
	}
}

// other test cases for French and sentences
func TestFrenchPalindrome(t *testing.T) {
	frenchPalindromes := []string{"été"}
	for _, w := range frenchPalindromes {
		if !IsPalindrome(w) {
			t.Errorf(`IsPalindrome(%q) = false`, w)
		}
	}
}

func TestSentencePalindrome(t *testing.T) {
	sentences := []string{"A man, a plan, a canal: Panama"}
	for _, s := range sentences {
		if !IsPalindrome(s) {
			t.Errorf(`IsPalindrome(%q) = false`, s)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"étre", false},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},
		{"desserts", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf(`IsPalindrome(%q) = %t`, test.input, got)
		}
	}
}
