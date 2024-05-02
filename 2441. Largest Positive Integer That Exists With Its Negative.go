// Time C.(nlogn), Space C. (1)
// runtime 15ms

func findMaxK(nums []int) int {

	slices.Sort(nums)

	i, j := 0, len(nums)-1

	for i < len(nums) && j > 0 {
		if nums[i]+nums[j] == 0 {
			return nums[j]
		} else if nums[i]*-1 > nums[j] {
			i++
		} else {
			j--
		}
	}

	return -1
}