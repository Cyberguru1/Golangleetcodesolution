// Free from hard ques
// obivously not my solution still learning the ropes

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solve(i, ringI int, ring, key string, dp *[][]int) int {
	if i >= len(key) {
		return 0
	}
	if (*dp)[i][ringI] != -1 {
		return (*dp)[i][ringI]
	}
	ans := math.MaxInt16
	newRingI := -1
	for I := ringI; I >= 0; I-- {
		if key[i] == ring[I] {
			newRingI = I
			ans = min(ans, 1+abs(ringI-I)+solve(i+1, newRingI, ring, key, dp))
			ans = min(ans, 1+len(ring)-abs(ringI-I)+solve(i+1, newRingI, ring, key, dp))
		}
	}
	for I := ringI; I < len(ring); I++ {
		if key[i] == ring[I] {
			newRingI = I
			ans = min(ans, 1+abs(ringI-I)+solve(i+1, newRingI, ring, key, dp))
			ans = min(ans, 1+len(ring)-abs(ringI-I)+solve(i+1, newRingI, ring, key, dp))
		}
	}
	(*dp)[i][ringI] = ans
	return ans
}
func Init2d[T any](n, m int, val T) [][]T {
	arr := make([][]T, n)
	for i := range n {
		arr[i] = make([]T, m)
	}
	for i := range n {
		for j := range m {
			arr[i][j] = val
		}
	}
	return arr
}

func findRotateSteps(ring string, key string) int {
	dp := Init2d(105, 105, -1)
	return solve(0, 0, ring, key, &dp)
}