func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	mod := 1000000007

	prev, curr := make([][]int, m+2), make([][]int, m+2)

	for i := range prev {
		prev[i] = make([]int, n+2)
		curr[i] = make([]int, n+2)
	}

	for i := 1; i <= m; i++ {
		prev[i][0], prev[i][n+1] = 1, 1
		curr[i][0], curr[i][n+1] = 1, 1
	}

	for j := 1; j <= n; j++ {
		prev[0][j], prev[m+1][j] = 1, 1
		curr[0][j], curr[m+1][j] = 1, 1
	}

	for move := 1; move <= maxMove; move++ {
		prev, curr = curr, prev

		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				leftR := (prev[i][j-1] + prev[i][j+1]) % mod
				upD := (prev[i-1][j] + prev[i+1][j]) % mod
				curr[i][j] = (leftR + upD) % mod
			}
		}
	}

	return curr[startRow+1][startColumn+1]

}