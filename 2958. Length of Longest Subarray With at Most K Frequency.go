// Time C. O(n), Space C. O(n)
// runtime 108ms

func maxSubarrayLength(nums []int, k int) (count int) {

	mapp := map[int]int{}

	l := 0
	for r, kk := range nums {
		mapp[kk]++

		for mapp[kk] > k {
			mapp[nums[l]]--
			l++
		}

		if r-l+1 > count {
			count = r - l + 1
		}
	}

	return
}
