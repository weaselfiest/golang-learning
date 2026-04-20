package bubblesort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "unsorted slice",
			input: []int{5, 3, 8, 1, 2},
			want:  []int{1, 2, 3, 5, 8},
		},
		{
			name:  "already sorted",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "reverse sorted",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "duplicates",
			input: []int{3, 1, 2, 1, 3},
			want:  []int{1, 1, 2, 3, 3},
		},
		{
			name:  "single element",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "negative numbers",
			input: []int{-3, 0, -1, 2, -5},
			want:  []int{-5, -3, -1, 0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.input)
			if !equal(tt.input, tt.want) {
				t.Errorf("got %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
