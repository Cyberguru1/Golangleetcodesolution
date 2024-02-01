// Time C. O(n), Space C. O(n+m), where m is size of arr
// runtime 183ms

import "slices"

func divideArray(nums []int, k int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)

	for i:=3; i < len(nums)+ 3; i = i+3{
		arr := []int{}
		if (nums[i-2] - nums[i-3] <= k) && (nums[i-1] - nums[i-2] <= k) && (nums[i-1] - nums[i-3] <= k){
			arr = append(arr, nums[i-3], nums[i-2], nums[i-1])
			result = append(result, arr)
			arr = arr[:0]
		}else {
			return [][]int{}
		}
	}
	return result
}