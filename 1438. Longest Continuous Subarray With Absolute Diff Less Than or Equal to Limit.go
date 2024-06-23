// using queue would have been better
// how about O(1) space

func longestSubarray(nums []int, limit int) (res int) {
	left, right, mn, mx := 0, 0, 0, 0

	for right < len(nums) {
		if nums[mn] >= nums[right] {
			mn = right
		}
		if nums[mx] <= nums[right] {
			mx = right
		}

		if nums[mx] - nums[mn] <= limit {
			right++
		} else {
			res = max(res, right - left)
			left = min(mn, mx) + 1
			mn, mx = left, left
			right = left + 1
		}
	}
	// to be sure
	res = max(res, right - left)
	return
}
