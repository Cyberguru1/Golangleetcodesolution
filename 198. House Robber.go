// Time C. O(n), Space C. O(n)
// runtime 1ms

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	memo := make([]int, len(nums)+1)

	memo[0] = 0
	memo[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		memo[i+1] = max(memo[i], memo[i-1]+val)
	}
	return memo[len(nums)]
}

func rob(nums []int) int {

	if len(nums) < 0 {
		return 0
	}

	prev1 := 0
	prev2 := 0

	for _, k := range nums {
		tmp := prev1
		prev1 = max(prev2+k, prev1)
		prev2 = tmp
	}

	return prev2
}