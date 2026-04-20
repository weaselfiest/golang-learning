package emailvalidator

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid simple",
			input: "user@example.com",
			want:  true,
		},
		{
			name:  "valid subdomain",
			input: "user@sub.example.com",
			want:  true,
		},
		{
			name:  "no dot after @",
			input: "user@example",
			want:  false,
		},
		{
			name:  "no local part",
			input: "@example.com",
			want:  false,
		},
		{
			name:  "dot immediately after @",
			input: "user@.com",
			want:  false,
		},
		{
			name:  "ends with dot",
			input: "user@example.",
			want:  false,
		},
		{
			name:  "space inside",
			input: "us er@example.com",
			want:  false,
		},
		{
			name:  "double @",
			input: "user@@example.com",
			want:  false,
		},
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
		{
			name:  "only @",
			input: "@",
			want:  false,
		},
		{
			name:  "no @",
			input: "userexample.com",
			want:  false,
		},
		{
			name:  "space at start",
			input: " user@example.com",
			want:  false,
		},
		{
			name:  "space at end",
			input: "user@example.com ",
			want:  false,
		},
		{
			name:  "valid with plus in local part",
			input: "user+tag@example.com",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidEmail(tt.input)
			if got != tt.want {
				t.Errorf("IsValidEmail(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
