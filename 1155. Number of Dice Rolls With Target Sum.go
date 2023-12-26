// Time C. O(n^k) Space C. O(n*target)
// medium chall 

func numRollsToTarget(n int, k int, target int) int {
	dp := make(map[int]map[int]int, n + 1)
	dp[0] = map[int]int{}
	dp[0][0] = 1
	mod := 1000000007
	for i := 1; i <= n; i++ {
		 dp[i] = map[int]int{}
		for j := 1; j <= k; j++ {
			for c := target; c >= j; c-- {
				if (dp[i - 1][c - j] != 0){
					dp[i][c] += dp[i - 1][c - j]
					dp[i][c] = dp[i][c] % mod
				}
			}
		}
	}

	return dp[n][target]
}
