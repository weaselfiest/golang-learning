package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"пустая строка", "", true},
		{"один символ", "a", true},
		{"простой палиндром", "racecar", true},
		{"не палиндром", "hello", false},
		{"с учётом регистра", "Racecar", true},
		{"кириллица", "Довод", true},
		{"русский палиндром с пробелами", "А роза упала на лапу азора", true},
		{"иероглифы", "日本語語本日", true},
		{"английский со знаками", "A man, a plan, a canal Panama", true},
		{"только пробелы и пунктуация", "...", true},
		{"почти палиндром", "racecars", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
