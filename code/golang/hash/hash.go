package hash

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	n := len(s)
	m := make(map[byte]byte)

	for i := 0; i < n; i++ {
		val1, ok1 := m[s[i]]
		val2, ok2 := m[t[i]]
		if !ok1 && !ok2 {
			m[s[i]] = t[i]
			m[t[i]] = s[i]
		} else if ok1 && ok2 && val1 == t[i] && val2 == s[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
