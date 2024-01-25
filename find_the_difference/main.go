package find_the_difference

func findTheDifference(s string, t string) byte {
	sumS, sumT := 0, 0

	for i := range s {
		sumS += int(s[i])
	}

	for i := range t {
		sumT += int(t[i])
	}

	return byte(sumT - sumS)
}
