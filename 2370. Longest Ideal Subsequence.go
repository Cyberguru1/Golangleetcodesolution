// Time C. O(n), Space C. O(1)
// runtime 28ms

func longestIdealString(s string, k int) int {

	dp := make([]int, 27)
	n := len(s)

	for i := n - 1; i >= 0; i-- {
		idx := int(s[i] - 'a')
		maxi := -1 << 31
		l := max(idx-k, 0)
		r := min(idx+k, 26)
		for j := l; j <= r; j++ {
			maxi = max(maxi, dp[j])
		}
		dp[idx] = maxi + 1
	}

	return slices.Max(dp)
}