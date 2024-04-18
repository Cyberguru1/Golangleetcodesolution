// Time C. O(n^2), Space C. O(1)
// Runtime 39ms

func islandPerimeter(grid [][]int) int {

	islands, neighbors := 0, 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				islands += 4
				// check left and right neighbors then increment if exists
				if i-1 > -1 && grid[i-1][j] == 1 {
					neighbors += 2
				}
				if j-1 > -1 && grid[i][j-1] == 1 {
					neighbors += 2
				}
			}
		}
	}
	return islands - neighbors
}