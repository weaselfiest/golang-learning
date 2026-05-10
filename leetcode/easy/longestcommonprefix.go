package easy

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 0; i < len(prefix); i++ {
		for _, str := range strs[1:] {
			if i >= len(str) || str[i] != prefix[i] {
				return prefix[:i]
			}
		}
	}

	return prefix
}
