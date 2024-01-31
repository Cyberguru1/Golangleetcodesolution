// Time C. O(n), Space C. O(n)
// runtime 146ms

func dailyTemperatures(temperatures []int) []int {
	mstack := make([]int, 0)
	st := make([]int, len(temperatures))

	for i, v := range temperatures {
		for len(mstack) > 0 && temperatures[mstack[len(mstack)-1]] < v {
			ind := mstack[len(mstack)-1]
			mstack = mstack[:len(mstack)-1]
			st[ind] = i - ind
		}
		mstack = append(mstack, i)

	}

	return st   
}