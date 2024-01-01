package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		chars := []rune(strs[i])
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		res[string(chars)] = append(res[string(chars)], strs[i])
	}

	var ans [][]string

	for _, v := range res {
		ans = append(ans, v)
	}
	return ans
}
