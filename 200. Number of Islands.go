// TIme C. O(M*N)/ Space C. O(1)
// Best Runtime 0ms

func numIslands(grid [][]byte) (count int ){
	
	var dfs func(grid [][]byte, r, c int)

	dfs = func(grid [][]byte, r, c int){
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid [r][c] != '1'{
			return
		}

		grid[r][c] = 'v' // to show we have visited
		// move on all direction
		dfs(grid, r + 1, c) // move down
		dfs(grid, r - 1, c) // move up
		dfs(grid, r, c + 1) // move right
		dfs(grid, r, c - 1) // move left
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return
}
