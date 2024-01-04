package main

func leastInterval(tasks []byte, n int) int {
	maxF := 0
	mapS := make(map[byte]int)

	for i := 0; i < len(tasks); i++ {
		mapS[tasks[i]]++
		if mapS[tasks[i]] > maxF {
			maxF = mapS[tasks[i]]
		}
	}

	res := (maxF - 1) * (n + 1)

	for _, v := range mapS {
		if v == maxF {
			res++
		}
	}

	res = max(res, len(tasks))
	return res
}
