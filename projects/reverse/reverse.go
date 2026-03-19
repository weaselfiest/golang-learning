package reverse

func reverse(s string) string {
	strRunes := []rune(s)
	for i, j := 0, len(strRunes)-1; i < j; i, j = i+1, j-1 {
		strRunes[i], strRunes[j] = strRunes[j], strRunes[i]
	}
	return string(strRunes)
}
