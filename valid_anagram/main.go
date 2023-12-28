package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[string]int)
	mapT := make(map[string]int)

	for i := 0; i < len(s); i++ {
		mapS[string(s[i])]++
		mapT[string(t[i])]++
	}

	for k := range mapS {
		if mapS[k] != mapT[k] {
			return false
		}
	}
	return true
}
