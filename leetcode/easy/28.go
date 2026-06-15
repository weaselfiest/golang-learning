package easy

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)

	for i := 0; i <= n-m; i++ {
		j := 0

		for j < m && haystack[i+j] == needle[j] {
			j++
		}

		if j == m {
			return i
		}
	}

	return -1
}
