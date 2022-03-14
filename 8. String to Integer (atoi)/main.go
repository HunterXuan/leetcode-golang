package main

func myAtoi(s string) int {
	started := false
	flag := 1
	num := 0
	for i := 0; i < len(s); i++ {
		if started {
			if s[i] <= '9' && s[i] >= '0' {
				num = num * 10 + int(s[i]) - '0'
				if num > 1 << 31 - 1 {
					break
				}
			} else {
				break
			}
		} else if s[i] == ' ' {
			continue
		} else if s[i] == '-' {
			flag = -1
			started = true
		} else if s[i] == '+' {
			flag = 1
			started = true
		} else if s[i] <= '9' && s[i] >= '0' {
			num = int(s[i]) - '0'
			started = true
		} else {
			break
		}
	}

	num = num * flag
	if num > 1 << 31 - 1 {
		num = 1 << 31 - 1
	} else if num < -(1 << 31) {
		num = -(1 << 31)
	}

	return num
}