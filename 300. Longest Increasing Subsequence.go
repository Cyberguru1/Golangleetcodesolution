// Time C. O(n^2), Space C. O(n)

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	for k, _ := range dp {
		dp[k] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

	}

	return slices.Max(dp)

}

