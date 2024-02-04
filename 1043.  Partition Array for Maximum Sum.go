

func maxSumAfterPartitioning(arr []int, k int) int {
	ln := len(arr)
	dp := make([]int,k)

	for i:=1;i<ln+1;i++{
		max_item := 0
		max_value := 0

		for j:=1;j<min(k,i)+1;j++{
			max_item = max(max_item, arr[i - j])
			max_value = max(max_value, dp[(i - j)%k] + max_item * j)
		}
		dp[i%k] = max_value
	}
	return dp[ln%k]
}