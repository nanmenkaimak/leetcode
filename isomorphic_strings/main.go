package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	map1 := make([]byte, 256)
	map2 := make([]byte, 256)

	for i := 0; i < len(s); i++ {
		if map1[s[i]] != map2[t[i]] {
			return false
		}
		map1[s[i]] = byte(i + 1)
		map2[t[i]] = byte(i + 1)
	}
	return true
}
