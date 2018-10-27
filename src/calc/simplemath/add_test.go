package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 3)
	if r != 4 {
		t.Error("Add(1,3) failed. Got %d, expected 4", r)
	}
}
