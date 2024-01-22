package backspace_string_compare

func backspaceCompare(s string, t string) bool {
	first, second := "", ""

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			first += string(s[i])
		} else if s[i] == '#' && len(first) > 0 {
			first = first[:len(first)-1]
		}
	}

	for i := 0; i < len(t); i++ {
		if t[i] >= 'a' && t[i] <= 'z' {
			second += string(t[i])
		} else if t[i] == '#' && len(second) > 0 {
			second = second[:len(second)-1]
		}
	}

	return first == second
}
