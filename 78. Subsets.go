// Time C. O(2^n), Space C. O(n)

func subsets(nums []int) (subsets [][]int) {
	var subbs func(res []int, ind int)
	subbs = func(res []int, ind int) {
		subsets = append(subsets, append([]int{}, res...))
		for j := ind; j < len(nums); j++ {
			res = append(res, nums[j])
			subbs(res, j+1)
			res = res[:len(res)-1]
		}
	}
	subbs([]int{}, 0)
	return
}
