//Time C. (n), Space C. (1)
//Runtime 0ms

func matrixScore(grid [][]int) (tsum int) {

	r, c := len(grid), len(grid[0])

	for i := 0; i < r; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < c; j++ {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}

	for i := 0; i < c; i++ {
		nz := 0
		for j := 0; j < r; j++ {
			if grid[j][i] == 0 {
				nz++
			}
		}
		if nz > (r - nz) {
			for j := 0; j < r; j++ {
				grid[j][i] = 1 - grid[j][i]
			}
		}
	}

	for i := 0; i < r; i++ {
		num := 0
		for j := 0; j < c; j++ {
			num += (grid[i][j] * (1 << (c - j - 1)))
		}
		tsum += num
	}
	return
}