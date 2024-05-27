// Time C. O(nlogn), Space C. O(1)
// Runtime 2ms

func specialArray(nums []int) int {

	sort.Slice(nums, func(i, j int) bool{
		return nums[j] > nums[i]
	})

	lent := len(nums)

	for lent > 0 {
		val := lent
		count := 0
		for _, k := range nums {
			if k >= val {
				count++
			}
		}
		if count == val {
			return val
		}
		lent--
	}

	return -1
}