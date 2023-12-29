// leetcode hard
// Time C. O(n*d) Space C. O(n)
// runtime 46ms

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	var findmin func(j, k int) int
	dp := make(map[[2]int]int)

	findmin = func(j, k int) int {

		if val, ok := dp[[2]int{j, k}]; ok {
			return val
		}
		if (j == n || k == 0) {
			if (j == n && k == 0) {
				return 0
			}else {
				return 99999
			}
		}

		result := 99999
		dayD := 0

		for i := j; i <= (n - k); i++ {
			dayD = max(dayD, jobDifficulty[i])
			result = min(result, (dayD + findmin(i + 1, k - 1)))
		}
		dp[[2]int{j, k}] = result
		return result
	}

	res := findmin(0, d)
	if res == 99999 {
		return -1
	}

	return res
}