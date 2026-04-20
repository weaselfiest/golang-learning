package wordcount

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "basic",
			input: "go is great and go is fun",
			want:  map[string]int{"go": 2, "is": 2, "great": 1, "and": 1, "fun": 1},
		},
		{
			name:  "case insensitive",
			input: "Go go GO",
			want:  map[string]int{"go": 3},
		},
		{
			name:  "single word",
			input: "hello",
			want:  map[string]int{"hello": 1},
		},
		{
			name:  "empty string",
			input: "",
			want:  map[string]int{},
		},
		{
			name:  "extra spaces",
			input: "  hello   world  ",
			want:  map[string]int{"hello": 1, "world": 1},
		},
		{
			name:  "all same words",
			input: "yes yes yes",
			want:  map[string]int{"yes": 3},
		},
		{
			name:  "mixed case multiple words",
			input: "Go RUST rust go Python python PYTHON",
			want:  map[string]int{"go": 2, "rust": 2, "python": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordCount(tt.input)
			if !equalMaps(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCountNotNil(t *testing.T) {
	got := WordCount("")
	if got == nil {
		t.Error("expected non-nil map for empty input, got nil")
	}
}

func equalMaps(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
