package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected string
	}{
		{"plain number", 1, "1"},
		{"divisible by 3", 3, "Fizz"},
		{"divisible by 5", 5, "Buzz"},
		{"divisible by 15", 15, "FizzBuzz"},
		{"divisible by 30", 30, "FizzBuzz"},
		{"not divisible", 7, "7"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fizzBuzz(tt.num)
			if result != tt.expected {
				t.Errorf("ожидали %s, получили %s", tt.expected, result)
			}
		})
	}
}
