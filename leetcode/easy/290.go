package easy

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	p2w := make(map[byte]string)
	w2p := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]

		if val, ok := p2w[p]; ok {
			if val != w {
				return false
			}
		} else {
			p2w[p] = w
		}

		if val, ok := w2p[w]; ok {
			if val != p {
				return false
			}
		} else {
			w2p[w] = p
		}
	}

	return true
}
