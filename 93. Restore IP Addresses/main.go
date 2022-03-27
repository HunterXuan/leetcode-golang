package main

import "fmt"

func restoreIpAddresses(s string) []string {
	sLen := len(s)
	result := make([]string, 0)
	for i := 1; i < 4; i ++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				if i+j+k < sLen {
					part1, part2, part3, part4 := s[0:i], s[i:i+j], s[i+j:i+j+k], s[i+j+k:]
					if validPart(part1) && validPart(part2) && validPart(part3) && validPart(part4) {
						result = append(result, fmt.Sprintf("%v.%v.%v.%v", part1, part2, part3, part4))
					}
				}
			}
		}
	}

	return result
}

func validPart(s string) bool {
	sLen := len(s)
	if sLen == 0 || sLen > 3 {
		return false
	}

	if sLen > 1 && s[0] == '0' {
		return false
	}

	num := 0
	for _, c := range s {
		num = num * 10 + int(c - '0')
	}

	if num > 255 {
		return false
	}

	return true
}