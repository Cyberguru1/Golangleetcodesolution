// Time C. O(N^2), Space C. O(n)
// Runtime 8ms

func largestLocal(grid [][]int) [][]int {
	result := make([][]int, 0)
	for row := range len(grid) - 2 {
		r := make([]int, 0) // row of resulting matrix
		for col := range len(grid[0]) - 2 {
			m := grid[row][col]
			for i := range 3 {
				for j := range 3 {
					m = max(m, grid[row+i][col+j])
				}
			}
			r = append(r, m)
		}
		result = append(result, r)
	}
	return result
}