package main

import "strconv"

func evalRPN(tokens []string) int {
	var operations []int

	for i := 0; i < len(tokens); i++ {
		res := 0
		num, err := strconv.Atoi(tokens[i])
		if err == nil {
			operations = append(operations, num)
		} else {
			number1 := operations[len(operations)-1]
			number2 := operations[len(operations)-2]
			operations = operations[:len(operations)-2]

			if tokens[i] == "+" {
				res = number1 + number2
			} else if tokens[i] == "-" {
				res = number1 - number2
			} else if tokens[i] == "*" {
				res = number1 * number2
			} else if tokens[i] == "/" {
				res = number1 / number2
			}
			operations = append(operations, res)
		}
	}
	return operations[0]
}
