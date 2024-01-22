package jewels_and_stones

func numJewelsInStones(jewels string, stones string) int {
	maps := make(map[byte]int)
	for i := 0; i < len(jewels); i++ {
		maps[jewels[i]]++
	}
	cnt := 0

	for i := 0; i < len(stones); i++ {
		if maps[stones[i]] > 0 {
			cnt++
		}
	}
	return cnt
}
