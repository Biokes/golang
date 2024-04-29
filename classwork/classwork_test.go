package main

import (
	"testing"
)

func TestAdditonMethod(t *testing.T) {
	num := add(5, 7)
	if num != 12 {
		t.Errorf("add(5, 7) = %d; want 12", num)
	}
}
