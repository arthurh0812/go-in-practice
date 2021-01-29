package ints

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small, non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 { // skip words that have no set integers
			continue
		}
		for j := 0; j < 64; j++ {
			// if the word at j is set
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements that are set.
func (s *IntSet) Len() (count int) {
	for _, w := range s.words {
		for ; w != 0; w = w >> 1 {
			if w&1 == 1 {
				count++
			}
		}
	}
	return
}

// Has reports whether this integer set also has includes the given non-negative integer x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds non-negative integer x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Remove removes non-negative integer x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		var shift uint64 = 1 << bit
		if s.words[word]&shift != 0 {
			s.words[word] ^= shift
		}
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems returns a slice containing the elements of the integer set
func (s *IntSet) Elems() (elements []int) {
	// that way, we know what capacity the slice has to have in the first place
	ln := s.Len()
	elements = make([]int, 0, ln)
	for i, word := range s.words {
		if word == 0 { // skip words that have no set integers
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(uint64(1<<j)) != 0 {
				elements = append(elements, i*64+j)
			}
		}
	}
	return elements
}

// Clear removes all the elements from a set.
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

// Copy returns a copy of this integer set.
// Copy allocates a new underlying array for the new integer set's words.
func (s *IntSet) Copy() *IntSet {
	newSet := &IntSet{}
	newSet.words = make([]uint64, 0, len(s.words))
	copy(newSet.words, s.words)

	return newSet
}
