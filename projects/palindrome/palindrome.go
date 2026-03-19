package palindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	strClear := make([]rune, 0, len(s))
	for _, el := range s {
		if unicode.IsLetter(el) {
			el = unicode.ToLower(el)
			strClear = append(strClear, el)
		}
	}

	for i, j := 0, len(strClear)-1; i < j; i, j = i+1, j-1 {
		if strClear[i] != strClear[j] {
			return false
		}
	}
	return true
}
