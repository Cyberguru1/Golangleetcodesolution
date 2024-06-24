// Time C. O(n), Space C. O(1)

func minKBitFlips(nums []int, k int) int {
	ans, flips := 0, 0
	for i := 0; i < len(nums); i++ {
		if (nums[i] + flips) % 2 == 0 {
			if i > len(nums) - k {
				return -1
			} else {
				ans++
				flips++
				nums[i] = -1
			}
		}
		if i + 1 >= k {
			if nums[i - k + 1] < 0 {
				flips--
			}
		}
	}
	return ans
}
