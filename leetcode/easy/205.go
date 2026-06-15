package easy

func isIsomorphic(s string, t string) bool {
	st := make(map[byte]byte)
	ts := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]

		if v, ok := st[c1]; ok {
			if v != c2 {
				return false
			}
		} else {
			st[c1] = c2
		}

		if v, ok := ts[c2]; ok {
			if v != c1 {
				return false
			}
		} else {
			ts[c2] = c1
		}
	}

	return true
}
