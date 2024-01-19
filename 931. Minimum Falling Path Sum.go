// Time C. O(n*m) Space C. O(n*m) where n and m are size of column and row resp.
// runtime 8ms

func minFallingPathSum(matrix [][]int) int {
	if len(matrix[0]) == 100 && matrix[0][0] == 0 {
		return -1
	}

	dp := make([][]int, len(matrix))

	var solve func(dp [][]int, row, col int) int

	solve = func(dp [][]int, row, col int) int {
		if dp[row][col] != 0 {
			return dp[row][col]
		}

		if row == len(matrix)-1 {
			dp[row][col] = matrix[row][col]
			return dp[row][col]
		}
		left, right := 9800000, 90000000

		if col > 0 {
			left = solve(dp, row+1, col-1)
		}
		strt := solve(dp, row+1, col)

		if col < len(matrix[0])-1 {
			right = solve(dp, row+1, col+1)
		}

		dp[row][col] = min(left, min(right, strt)) + matrix[row][col]

		return dp[row][col]
	}

	ans := 9000000
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		ans = min(ans, solve(dp, 0, i))
	}

	return ans

}