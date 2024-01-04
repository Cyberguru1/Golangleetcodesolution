// first brute force solution
// Time C. O(n), Space C. O(n)

func minOperations(nums []int) int {

	result := 0
	dibs := make(map[int]int)

	for _, v := range nums {
		dibs[v]++
	}

	for _, v := range dibs {

		if v == 1 {
			return -1
		} else if v%3 == 0 {
			result += v / 3
		} else if v%3 == 1 {
			result += ((v / 3) - 1) + 2
		} else if v%3 == 2 {
			result += (v / 3) + 1
		} else if v%2 == 0 {
			result += v / 2
		} else {
			return -1
		}
	}

	return result

}