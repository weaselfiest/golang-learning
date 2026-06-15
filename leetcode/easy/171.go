package easy

func titleToNumber(columnTitle string) int {
	result := 0

	for _, ch := range columnTitle {
		value := int(ch-'A') + 1
		result = result*26 + value
	}

	return result
}
