package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			dp[i] = dp[i-1]
		}

		if i > 0 {
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
				if i > 1 {
					dp[i] = dp[i] + dp[i-2]
				} else {
					dp[i] = dp[i] + 1
				}
			}
		}

		if dp[i] == 0 {
			return 0
		}
	}

	return dp[len(dp) - 1]
}