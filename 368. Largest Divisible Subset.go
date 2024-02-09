// first solution
// Time C. O(n^2), Space C. O(n)
// runtime 159ms

func largestDivisibleSubset(nums []int) []int {

	maxx := 0
	maxl := []int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	check := func(arr []int) (res bool) {
		res = true
		for i, k := range arr {
			for j, h := range arr {
				if i == j {
					continue
				}
				if k%h == 0 || h%k == 0 {
					res = res && true
				} else {
					res = res && false
					return
				}
			}
		}
		return
	}

	for i, l := range nums {

		dp := []int{}
		dp = append(dp, l)
		lums := append([]int{}, nums[i+1:]...)
		lums = append(lums, nums[:i]...)

		for _, k := range lums {
			dp = append(dp, k)

			if !check(dp) {
				dp = dp[:len(dp)-1]
			}
		}
		if len(dp) > maxx {
			maxx = len(dp)
			maxl = dp
		}
	}
	return maxl
}

// second solution dynamic approach