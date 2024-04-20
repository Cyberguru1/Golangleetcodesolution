// comming back for this

func findFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
	var answer [][]int

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if land[r][c] == 1 {
				rMax, cMax := -1, -1
				dfs(land, r, c, &rMax, &cMax)
				answer = append(answer, []int{r, c, rMax, cMax})
			}
		}
	}

	return answer
}

func dfs(land [][]int, r, c int, rMax, cMax *int) {
	if r < 0 || r >= len(land) || c < 0 || c >= len(land[0]) || land[r][c] == 0 {
		return
	}

	// Mark current cell as visited
	land[r][c] = 0

	// Update rMax and cMax if current cell is more 'bottom-right'
	if r > *rMax {
		*rMax = r
	}
	if c > *cMax {
		*cMax = c
	}

	// Explore the cell to the right (c + 1) and below (r + 1)
	dfs(land, r, c+1, rMax, cMax) // Move right
	dfs(land, r+1, c, rMax, cMax) // Move down
}
