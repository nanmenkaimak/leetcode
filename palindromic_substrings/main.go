package main

func countSubstrings(s string) int {
	res := 0
	maps := make(map[byte][]int)

	for i := 0; i < len(s); i++ {
		maps[s[i]] = append(maps[s[i]], i)

		if len(maps[s[i]]) > 1 {
			for _, v := range maps[s[i]] {
				if v != i && isPalindrom(s[v:i+1]) {
					res++
				}
			}
		}
	}
	return res + len(s)
}

func isPalindrom(s string) bool {
	if len(s) == 1 || len(s) == 0 {
		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
