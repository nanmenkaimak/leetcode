package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	idx := 0

	for i := 0; i < len(t); i++ {
		if t[i] == s[idx] {
			idx++
		}
		if idx == len(s) {
			return true
		}
	}

	return false
}
