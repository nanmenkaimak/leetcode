package main

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	maps := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		maps[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		_, ok := maps[nums2[i]]
		if ok {
			maps[nums2[i]]--
			res = append(res, nums2[i])
			if maps[nums2[i]] == 0 {
				delete(maps, nums2[i])
			}
		}
	}
	return res
}
