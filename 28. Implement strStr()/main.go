package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	haystackLen, needleLen := len(haystack), len(needle)
	if needleLen > haystackLen {
		return -1
	}

	for i := 0; i < haystackLen - needleLen + 1; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}