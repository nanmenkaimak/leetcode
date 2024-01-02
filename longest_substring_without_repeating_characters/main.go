package main

func lengthOfLongestSubstring(s string) int {
	res := 0
	mapS := make(map[byte]int)

	left := 0

	for i := 0; i < len(s); i++ {
		mapS[s[i]]++
		for mapS[s[i]] > 1 {
			mapS[s[left]]--
			left++
		}

		res = max(res, i-left+1)
	}
	return res
}
