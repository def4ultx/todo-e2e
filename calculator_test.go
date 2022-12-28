package main

import "testing"

func TestPlus(t *testing.T) {
	actual := Plus(1, 2)
	expected := 3
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestMinus(t *testing.T) {
	actual := Minus(1, 2)
	expected := -1
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
