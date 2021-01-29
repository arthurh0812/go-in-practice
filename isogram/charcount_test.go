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
