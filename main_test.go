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

// yoko
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

// tate
func TestCheck02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 0, "x")
	b.put(1, 1, "x")
	b.put(1, 2, "x")
	if b.check() != "yes" {
		t.Errorf("Mistake")
	}
}

// nashi
func TestCheck03(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "x")
	b.put(0, 1, "o")
	b.put(0, 2, "x")
	if b.check() != "no" {
		t.Errorf("Mistake")
	}
}

// naname
func TestCheck04(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 2, "x")
	b.put(1, 1, "x")
	b.put(2, 0, "x")
	if b.check() != "yes" {
		t.Errorf("Mistake")
	}
}
