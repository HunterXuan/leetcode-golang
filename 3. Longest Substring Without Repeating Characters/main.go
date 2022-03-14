package main

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	startIdx := 0
	strPosMap := make(map[int32]int)
	for idx, char := range s {
		pos, ok := strPosMap[char]
		strPosMap[char] = idx
		if ok && pos >= startIdx {
			startIdx = pos + 1
		} else {
			maxLen = max(idx - startIdx + 1, maxLen)
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}