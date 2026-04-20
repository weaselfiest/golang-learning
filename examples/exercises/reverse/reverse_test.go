package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected string
	}{
		{"hello", "hello", "olleh"},
		{"Привет", "Привет", "тевирП"},
		{"a", "a", "a"},
		{"", "", ""},
		{"日本語", "日本語", "語本日"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverse(tt.str)
			if result != tt.expected {
				t.Errorf("ожидали %s, получили %s", tt.expected, result)
			}
		})
	}
}
