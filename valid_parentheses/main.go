package main

var Parentheses = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, bracket := range s {
		par, ok := Parentheses[bracket]
		if ok {
			stack = append(stack, par)
			continue
		}

		if len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] != bracket {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
