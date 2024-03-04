// Time C. O(n), Space C. O(1)

func minDifference(nums []int) (res int) {
	lent := len(nums)
	if lent <= 4 { return }

	sort.Ints(nums)
	res = int(^uint(0) >> 1)

	l, r := 3, 0

	for l >= 0 && r <= 3 {
		res = min(max((nums[l]-nums[lent-1-r]), (nums[l]-nums[lent-1-r])*-1), res)
		r++
		l--
	}

	return
}
