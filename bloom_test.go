package main

import (
	"testing"
)

func TestBloom(t *testing.T) {
	bl := newBloom()
	bl = bl.Insert("hello")

	if !bl.IsPresent("hello") {
		t.Errorf("Insertion test failed")
	}

	if bl.IsPresent("world") {
		t.Errorf("function fidelity test failed")
	}
}
