// Path with maximum Gold
// Time C. O(4^n), Space C. O(n)

func getMaximumGold(grid [][]int) (res int) {

	r, c := len(grid), len(grid[0])

	var backtrack func(grid [][]int, i, j, curr int)

	backtrack = func(grid [][]int, i, j, curr int){

		if i >= r || j >= c || i < 0 || j < 0 || grid[i][j] == 0 {
			return
		}

		curr += grid[i][j]
		val := grid[i][j]

		grid[i][j] = 0

		res = max(res, curr)

		backtrack(grid, i-1, j, curr)
		backtrack(grid, i+1, j, curr)
		backtrack(grid, i, j-1, curr)
		backtrack(grid, i, j+1, curr)

		grid[i][j] = val
		curr -= val

	}

	for i := 0; i < r; i++{
		for j:= 0; j < c; j++ {
			if grid[i][j] != 0 {
				backtrack(grid, i, j, 0)
			}
		}
	}

	return
}