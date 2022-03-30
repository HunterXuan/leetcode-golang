package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len + s2Len != s3Len {
		return false
	}

	dp := make([][]bool, s3Len + 1)
	for idx, _ := range dp {
		dp[idx] = make([]bool, s1Len + 1)
	}
	dp[0][0] = true

	for i := 1; i <= s3Len; i++ {
		for j := 0; j <= i; j++ {
			k := i - j
			if j > s1Len || k > s2Len {
				continue
			}

			dp[i][j] = (j >= 1 && dp[i-1][j-1] && s3[i-1] == s1[j-1]) || (k >= 1 && dp[i-1][j] && s3[i-1] == s2[k-1])
		}
	}

	return dp[s3Len][s1Len]
}
