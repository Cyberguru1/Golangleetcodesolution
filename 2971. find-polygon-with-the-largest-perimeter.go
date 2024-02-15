// Time C. O(nlogn), Space C. O(1)
// runtime 136ms

func largestPerimeter(nums []int) int64 {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var res int64 = -1

	var prevS int64 = 0

	for i := 0; i < len(nums); i++ {
		if i >= 2 && int64(nums[i]) < prevS {
			res = int64(nums[i]) + prevS
		}

		prevS += int64(nums[i])
	}

	return res
}