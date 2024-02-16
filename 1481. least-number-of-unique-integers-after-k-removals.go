// Time C. O(nlogn) due to sorting, space complexity O(n)
// runtime 86ms
least-number-of-unique-integers-after-k-removals
func findLeastNumOfUniqueInts(arr []int, k int) int {
	mapp := map[int]int{}

	for _, j := range arr {
		mapp[j]++
	}

	freq := make([]int, len(mapp))
	c := 0
	for _, k := range mapp {
		freq[c] = k
		c++
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i] < freq[j]
	})

	rem := len(freq)

	for i := 0; i < len(freq) && k > 0; i++ {

		if freq[i] <= k {
			k -= freq[i]
			rem--
		}
	}

	return rem
}