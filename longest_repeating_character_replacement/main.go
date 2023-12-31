package main

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res, maxf, left := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]] += 1

		maxf = max(maxf, count[s[right]])

		if (right-left+1)-maxf > k {
			count[s[left]] -= 1
			left += 1
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
