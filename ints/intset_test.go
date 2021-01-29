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

	elems := []int{9, 12, 144}
	for _, e := range elems {
		if builtin[e] != true {
			t.Errorf("builtin[%d] = false", e)
		}
	}

	if ln := len(builtin); ln != 3 {
		t.Errorf("len(builtin) = %d, want 3", ln)
	}

	// should no be set
	x := 10
	if v, has := builtin[x]; has != false || v != false {
		t.Errorf("builtin[%d] = true", x)
	}

	// should be set
	y := 12
	delete(builtin, y)
	if v, has := builtin[y]; has != false || v != false {
		t.Errorf("builtin[%d] = true", y)
	}

	if ln := len(builtin); ln != 2 {
		t.Errorf("len(builtin) = %d, want 2", ln)
	}

	z := 144
	builtin = make(map[int]bool)
	if v, has := builtin[z]; has != false || v != false {
		t.Errorf("builtin[%d] = true", z)
	}

	intset.Add(9)
	intset.Add(12)
	intset.Add(144)

	for i, e := range intset.Elems() {
		if elems[i] != e {
			t.Errorf("intset[%d] != %d", e, elems[i])
		}
	}

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

	intset.Clear()

	if ln := intset.Len(); ln != 0 {
		t.Errorf("len(intset) = %d, want 0", ln)
	}

	if has := intset.Has(z); has != false {
		t.Errorf("intset[%d] = true", z)
	}
}
