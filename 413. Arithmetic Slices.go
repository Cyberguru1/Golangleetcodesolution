// Time C. O(n), Space C. O(1)
// runtime 0ms

func numberOfArithmeticSlices(nums []int) int {

	s, r := 0, 0

	for i := 0; i < len(nums)-2; i++ {
		if (nums[i+1] - nums[i]) == (nums[i+2] - nums[i+1]) {
			s++
		} else {
			r += s * (s + 1) / 2
			s = 0
		}
	}

	return r + s*(s+1)/2

}