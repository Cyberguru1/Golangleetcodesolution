// Time C. O(n), Space C. O(1)
// runtime 15ms

func missingNumber(nums []int) int {

	res := 0

	for i := 0; i < len(nums)+1; i++ {
		if i < len(nums) {
			res = nums[i] ^ res ^ i
		} else {
			res = res ^ i
		}
	}

	return res

}