
// third solution, space O(n) time O(N^2)
// runtime 187ms

func onesMinusZeros(grid [][]int) [][]int {
	n_rows := len(grid)
	n_cols := len(grid[0])

	sumRows := make([]int, n_rows)
	sumCols := make([]int, n_cols)

	for i := 0; i < n_rows; i++ {
		for j := 0; j < n_cols; j++ {

			if grid[i][j] == 1 {
				sumRows[i]++
			} else {
				sumRows[i]--
			}
		}
	}

	for i := 0; i < n_cols; i++ {
		for j := 0; j < n_rows; j++ {
			if grid[j][i] == 1 {
				sumCols[i]++
			} else {
				sumCols[i]--
			}
		}
	}

	for i, k := range grid {
		d := sumRows[i]
		for j, _ := range k {
			k[j] = d + sumCols[j]
		}
	}

	return grid
}

