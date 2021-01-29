package ints

import (
	"testing"
)

func TestIntSet(t *testing.T) {
	builtin := make(map[int]bool)
	intset := IntSet{}

	builtin[9] = true
	builtin[12] = true
	builtin[144] = true

	if ln := len(builtin); ln != 3 {
		t.Errorf("len(builtin) = %d, want 3", ln)
	}

	// should no be set
	x := 10
	if _, has := builtin[x]; has != false {
		t.Errorf("builtin[%d] = true", x)
	}

	// should be set
	y := 12
	delete(builtin, y)
	if _, has := builtin[y]; has != false {
		t.Errorf("builtin[%d] = true", y)
	}

	if ln := len(builtin); ln != 2 {
		t.Errorf("len(builtin) = %d, want 2", ln)
	}

	intset.Add(9)
	intset.Add(12)
	intset.Add(144)

	if ln := intset.Len(); ln != 3 {
		t.Errorf("len(intset) = %d, want 3", ln)
	}

	if has := intset.Has(x); has != false {
		t.Errorf("intset[%d] = true", x)
	}

	intset.Remove(y)
	if has := intset.Has(y); has != false {
		t.Errorf("intset[%d] = true", y)
	}

	if ln := intset.Len(); ln != 2 {
		t.Errorf("len(intset) = %d, want 2", ln)
	}
}
