package easy

import "unicode"

func isPalindrome1(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && !unicode.IsLetter(rune(s[left])) &&
			!unicode.IsDigit(rune(s[left])) {
			left++
		}

		for left < right && !unicode.IsLetter(rune(s[right])) &&
			!unicode.IsDigit(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) !=
			unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}
