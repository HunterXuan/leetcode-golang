package main

func longestPalindrome(s string) string {
	maxLen, maxStr := 0, ""

	for i := 0; i < len(s); i++ {
		tmp := s[i-maxLen:i+1]
		if checkPalindrome(tmp) {
			maxLen = maxLen + 1
			maxStr = tmp
		} else if i > maxLen {
			tmp = s[i-maxLen-1:i+1]
			if checkPalindrome(tmp) {
				maxLen = maxLen + 2
				maxStr = tmp
			}
		}
	}

	return maxStr
}

func checkPalindrome(s string) bool {
	strLen := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] != s[strLen - i - 1] {
			return false
		}
	}

	return true
}