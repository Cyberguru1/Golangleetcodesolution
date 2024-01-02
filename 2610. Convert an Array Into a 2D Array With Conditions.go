// first brute force solution
// Time C. O(nm) where n is size of nums and m is the max frequency of a number existing in nums
// Space C. O(n)
// Best runtime 9ms

func findMatrix(nums []int) [][]int {
	maxx := 0

	dubs := make(map[int]int)

	for _, k := range nums {
		dubs[k]++
		maxx = max(dubs[k], maxx)
	}

	result := make([][]int, maxx)

	for k, v := range dubs {
		for i := v - 1; i >= 0; i-- {
			result[i] = append(result[i], k)
		}
	}

	return result
}

// second solution (more intuitive)
// Time C. O(n), Space C. O(n)
// Best runtime 0ms

func findMatrix(nums []int) [][]int {
	dubs := make(map[int]int)

	result := make([][]int, 1)

	for _, k := range nums {

		if dubs[k] >= len(result) {
			result = append(result, []int{})
			result[dubs[k]] = append(result[dubs[k]], k)
		} else {
			result[dubs[k]] = append(result[dubs[k]], k)
		}

		dubs[k]++
	}

	return result
}
