// clealy not my solution


func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = 101
		}1531. String Compression II

	}

	var dfs func(int, int) int
	dfs = func(i, k int) int {
		if i+k >= n {
			return 0
		}
		if k < 0 {
			return 101
		}
		if dp[i][k] != 101 {
			return dp[i][k]
		}

		res := dfs(i+1, k-1) // Skip this character
		diff, same, length := 0, 0, 0

		// For all continuous s[i] characters (can skip k characters)
		for j := i; j < n; j++ {
			if k-diff < 0 {
				break
			}
			if s[i] == s[j] {
				same++
				if same <= 2 || same == 10 || same == 100 {
					length++
				}
			} else {
				diff++
			}
			res = int(math.Min(float64(res), float64(length+dfs(j+1, k-diff))))
		}

		dp[i][k] = res
		return res
	}

	return dfs(0, k)
}