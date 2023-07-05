package main

import (
	"testing"
)

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("Mistake")
	}
}

func TestPutToken02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 2, "x")
	if b.get(1, 2) != "x" {
		t.Errorf("Mistake")
	}
}

func TestCheck01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	b.put(1, 0, "o")
	b.put(2, 0, "o")
	if b.check() != "yes" {
		t.Errorf("Mistake")
	}
}
