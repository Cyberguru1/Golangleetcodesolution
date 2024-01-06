// dp problem
// runtime 108 ms

func jobScheduling(startTime []int, endTime []int, profit []int) int {

	var bisec func(arr [][]int, t int) int

	bisec = func(arr [][]int, t int) int {
		low, high := 0, len(arr)

		for low < high {
			mid := (low + high) / 2
			if arr[mid][0] > t {
				high = mid
			} else {
				low = mid + 1
			}
		}

		return low
	}

	n := len(startTime)
	job := make([][]int, n)
	dp := [][]int{{0, 0}}

	for i := 0; i < n; i++ {
		job[i] = []int{startTime[i], endTime[i], profit[i]}
	}

	sort.Slice(job, func(i, j int) bool {
		return job[i][1] < job[j][1]
	})

	for _, j := range job {
		start, end, prof := j[0], j[1], j[2]
		i := bisec(dp, start)

		nProf := dp[i-1][1] + prof

		if nProf > dp[len(dp)-1][1] {
			dp = append(dp, []int{end, nProf})
		}
	}

	return dp[len(dp)-1][1]

}