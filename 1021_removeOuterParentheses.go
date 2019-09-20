package random

func removeOuterParentheses(S string) string {
	stack := make([]rune, 0, len(S))
	var out string
	for _, b := range S {
		if len(stack) == 0 {
			stack = append(stack, b)
			continue
		}

		// if you get a closing the pull off the stack
		if b == ')' && stack[len(stack)-1] == '(' {
			// there's more than one left on the stack so keep tracking
			if len(stack) > 1 {
				out = out + string(b)
			}
			stack = stack[:len(stack)-1]
		} else {
			out = out + string(b)
			stack = append(stack, b)
		}
	}
	return out
}
