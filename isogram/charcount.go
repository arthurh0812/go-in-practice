package isogram

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

// ErrCharInvalid is the error returned when an invalid rune was encountered
var ErrCharInvalid = errors.New("invalid utf8 code point character")

// CharCount uses the given reader to count all utf8 encoded characters from its source
func CharCount(input io.Reader) (charCount map[rune]int, sizeCount map[int]int, err error) {
	in := bufio.NewReader(input)

	counts := make(map[rune]int)
	sizes := make(map[int]int)
	invalid := 0
	for {
		r, size, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		if r == unicode.ReplacementChar && size == 1 {
			invalid++
			continue
		}
		// increment the corresponding character and utf8 character length (size) in the maps
		counts[r]++
		sizes[size]++
	}

	if invalid != 0 {
		return counts, sizes, ErrCharInvalid
	}

	return counts, sizes, nil
}
