package main

func checkInclusion(s1 string, s2 string) bool {
	maps := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		maps[s1[i]-'a']++
	}

	start, end := 0, 0
	maps2 := make([]int, 26)
	for end < len(s2) {
		maps2[s2[end]-'a']++

		if end-start+1 == len(s1) {
			ok := compare(maps, maps2)
			if ok {
				return true
			}
		}

		if end-start+1 < len(s1) {
			end++
		} else {
			maps2[s2[start]-'a']--
			start++
			end++
		}
	}
	return false
}

func compare(map1, map2 []int) bool {
	for i := 0; i < 26; i++ {
		if map1[i] != map2[i] {
			return false
		}
	}
	return true
}
