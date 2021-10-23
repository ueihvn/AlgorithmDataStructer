package main

func isValidParentheses(s string) bool {
	stack := []string{}
	for _, v := range s {
		strv := string(v)
		if strv == "(" || strv == "[" || strv == "{" {
			stack = append(stack, string(v))
		} else if strv == ")" && stack[len(stack)-1] == "(" {
			stack = stack[0 : len(stack)-1]
		} else if strv == "]" && stack[len(stack)-1] == "[" {
			stack = stack[0 : len(stack)-1]
		} else if strv == "}" && stack[len(stack)-1] == "{" {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
