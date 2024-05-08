// Time C. O(NlogN), Space C. O(1)
// Runtime 20ms

func maximumHappinessSum(happiness []int, k int) (res int64) {

	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	r := 0
	for j := 0; j < k && j < len(happiness); j++ {
		rap := happiness[j] - r
		if rap < 0 {
			rap = 0
		}
		res += int64(rap)
		r++
	}
	return
}