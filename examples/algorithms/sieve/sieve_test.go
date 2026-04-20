package sieve

import (
	"slices"
	"testing"
)

func TestSieve(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected []int
	}{
		{"0", 0, []int{}},
		{"1", 1, []int{}},
		{"2", 2, []int{2}},
		{"10", 10, []int{2, 3, 5, 7}},
		{"9", 9, []int{2, 3, 5, 7}},
		{"11", 11, []int{2, 3, 5, 7, 11}},
		{"-5", -5, []int{}},
		{"100", 100,
			[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sieve(tt.num)
			if !slices.Equal(result, tt.expected) {
				t.Error(tt.expected, result)
			}
		})
	}
}
