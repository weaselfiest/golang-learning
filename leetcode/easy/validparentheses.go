package easy

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		bracket := s[i]
		switch bracket {
		case '(', '[', '{':
			stack = append(stack, bracket)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != pairs[bracket] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
