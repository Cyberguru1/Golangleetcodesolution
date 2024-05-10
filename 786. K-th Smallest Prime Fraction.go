// Time C. O(NlogN), Space C. O(N)
// Runtime 313ms

type cont struct {
	vals []int
	res  float64
}

func kthSmallestPrimeFraction(arr []int, k int) (res []int) {

	mapp := []cont{}

	for i, val := range arr {
		for k := i + 1; k < len(arr); k++ {
			mapp = append(mapp, cont{[]int{val, arr[k]}, float64(val) / float64(arr[k])})
		}
	}

	sort.Slice(mapp, func(i, j int) bool {
		return mapp[i].res < mapp[j].res
	})

	res = mapp[k-1].vals

	return
}