package leet

func IsSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	sl := len(s)
	tl := len(t)
	i := 0
	j := 0

	for i < sl || j < tl {
		if i == sl {
			break
		}
		if j == tl {
			break
		}
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == sl {
		return true
	}
	return false
}
