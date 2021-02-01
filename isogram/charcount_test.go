package isogram

import (
	"io"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		input     io.Reader
		char      rune
		charCount int
		size      int
		sizeCount int
		err       bool
	}{
		{
			input:     strings.NewReader("abcdefghijklmnopqrstuvwxyzq"),
			char:      'q',
			charCount: 2,
			size:      2,
			// meaning that there are 0 characters with a size of 2 bytes
			sizeCount: 0,
			err:       false,
		},
		{
			input:     strings.NewReader("the horses in the garden"),
			char:      'e',
			charCount: 4,
			size:      1,
			// meaning that there are 24 characters with a size of 1 byte
			sizeCount: 24,
			err:       false,
		},
		{
			input:     strings.NewReader("中文字文"),
			char:      '文',
			charCount: 2,
			size:      3,
			// meaning that there are 4 characters with a size of 3 bytes
			sizeCount: 4,
			err:       false,
		},
	}

	for _, test := range tests {
		counts, sizes, err := CharCount(test.input)
		// test if there should be an error, but in reality isn't
		if test.err && err == nil {
			t.Errorf(`CharCount().err = nil`)
		}
		if got := counts[test.char]; got != test.charCount {
			t.Errorf(`CharCount().counts[%v] = %d, want %d`, test.char, got, test.charCount)
		}
		if got := sizes[test.size]; got != test.sizeCount {
			t.Errorf(`CharCount().sizes[%d] = %d, want %d`, test.size, got, test.sizeCount)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		s      string
		sep    string
		pieces []string
	}{
		{
			s:      "a:b:c",
			sep:    ":",
			pieces: []string{"a", "b", "c"},
		},
		{
			s:      "a,long,walk",
			sep:    ",",
			pieces: []string{"a", "long", "walk"},
		},
		{
			s:      "on a friday night...",
			sep:    ".",
			pieces: []string{"on a friday night", "", "", ""},
		},
		{
			s:      "i like to move it!",
			sep:    ",",
			pieces: []string{"i like to move it!"},
		},
		{
			s:      "domain/myurl/site",
			sep:    "/",
			pieces: []string{"domain", "myurl", "site"},
		},
	}

	for _, test := range tests {
		pieces := strings.Split(test.s, test.sep)
		if got, want := len(pieces), len(test.pieces); got != want {
			t.Errorf("Split(%q, %q) returns words with length = %d, want %d",
				test.s, test.sep, got, want)
		}
	}
}
