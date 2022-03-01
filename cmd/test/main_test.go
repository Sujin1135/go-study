package main

import "testing"

func TestSquare(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81 but square(9) returns %d\n", rst)
	}
}
