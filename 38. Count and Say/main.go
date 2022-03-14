package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	result := countAndSayString("1")
	for i := 2; i < n; i++ {
		result = countAndSayString(result)
	}

	return result
}

func countAndSayString(s string) string {
	var result string

	curr, count := s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] == curr {
			count = count + 1
		} else {
			result = result + fmt.Sprintf("%v", count) + string(curr)
			curr, count = s[i], 1
		}
	}

	result = result + fmt.Sprintf("%v", count) + string(curr)

	return result
}