package main

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := [][]int{}
	for range m + 1 {
		dp = append(dp, make([]int, n+1))
	}
	for i := range m + 1 {
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
