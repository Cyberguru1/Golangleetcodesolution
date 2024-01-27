// Time C. O(n*k) Space C. O(k)
// runtime 8ms

func kInversePairs(n int, k int) int {
	mod := 1000000007
	dp := make([]int, k+1)
	dp[0] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[j] = (dp[j] + dp[j-1]) % mod
		}
		for j := k; j >= i; j-- {
			dp[j] = (dp[j] - dp[j-i] + mod) % mod
		}
	}

	return dp[k]

}