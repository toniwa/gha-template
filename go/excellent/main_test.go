package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{2, "even"},
		{3, "odd"},
		{0, "even"},
		{-1, "odd"},
		{-2, "even"},
	}

	for _, test := range tests {
		result := EvenOrOdd(test.number)
		if result != test.expected {
			t.Errorf("EvenOrOdd(%d) = %s; want %s", test.number, result, test.expected)
		}
	}
}
