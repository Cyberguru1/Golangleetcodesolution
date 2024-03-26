// Time C. O(n), Space C. O(1)
// Runtime 67ms

type void struct{}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func firstMissingPositive(nums []int) int {
	lent := len(nums)
	var mem void
	mapp := map[int]void{}
	for _, k := range nums {
		mapp[k] = mem
	}
	for i := 1; i <= lent*2; i++ {
		if _, ok := mapp[abs(i)]; !ok {
			return i
		}
	}
	return 0
}
41. First Missing Positive