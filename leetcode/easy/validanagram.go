package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	lettersCount := make(map[rune]int)

	for _, r := range s {
		lettersCount[r]++
	}

	for _, r := range t {
		lettersCount[r]--
		if lettersCount[r] < 0 {
			return false
		}
	}

	return true
}
