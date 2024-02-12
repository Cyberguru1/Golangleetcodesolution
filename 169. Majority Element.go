// Time C. O(nlogn), Space C. O(1)
// runtime 11ms

func majorityElement(nums []int) int {

	n := len(nums)

	if n == 1 {
		return nums[0]
	}
	n = (n / 2) + 1

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	c := 1
	init := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == init {
			c++
		} else {
			init = nums[i]
			c = 1
		}

		if c == n {
			return nums[i]
		}
	}
	return 0
}