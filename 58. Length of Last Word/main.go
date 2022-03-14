package main

func lengthOfLastWord(s string) int {
	start, end := -1, -1
	for idx := len(s) - 1; idx >= 0; idx-- {
		if end < 0 {
			if s[idx] == ' ' {
				continue
			}

			end = idx
		} else {
			if s[idx] == ' ' {
				start = idx
				break
			}
		}
	}

	return end - start
}