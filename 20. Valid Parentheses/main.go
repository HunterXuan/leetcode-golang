package main

func isValid(s string) bool {
	var stack []int32

	pair := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}

		if p, ok := pair[c]; ok {
			stackLen := len(stack)
			if stackLen > 0 && p == stack[stackLen - 1] {
				stack = stack[:stackLen-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}