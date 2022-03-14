package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	commonPrefix := findCommonPrefix(strs[0], strs[1])
	if commonPrefix == "" {
		return commonPrefix
	}

	for i := 2; i < len(strs); i++ {
		commonPrefix = findCommonPrefix(commonPrefix, strs[i])
		if commonPrefix == "" {
			return commonPrefix
		}
	}

	return commonPrefix
}

func findCommonPrefix(str1, str2 string) string {
	len1, len2 := len(str1), len(str2)
	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	if minLen == 0 {
		return ""
	}

	maxIdx := -1
	for i := 0; i < minLen; i++ {
		if str1[i] == str2[i] {
			maxIdx = i
		} else {
			break
		}
	}

	if maxIdx >= 0 {
		return str1[0:maxIdx+1]
	}

	return ""
}