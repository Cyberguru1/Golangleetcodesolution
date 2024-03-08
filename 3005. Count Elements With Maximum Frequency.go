// Time C. O(n), Space C. O(n)
// Best runtime 0ms

func maxFrequencyElements(nums []int) (res int) {

	mapp := map[int]int{}
	maxF := 0

	for _, v := range nums {
		mapp[v]++
		maxF = max(maxF, mapp[v])
	}

	for _, k := range mapp {
		if k == maxF {
			res += k
		}
	}
	return
}
