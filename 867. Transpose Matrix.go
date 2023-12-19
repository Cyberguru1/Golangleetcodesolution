// first solution, Time C. O(m*n), Space C. O(m*n)
// runtime record 9ms

func transpose(matrix [][]int) [][]int {
	lr := len(matrix)
	lc := len(matrix[0])
	new_mat := make([][]int, lc)

	for i:= 0; i < lc; i++ {
		new_mat[i] = make([]int, lr)
		for j:= 0; j < lr; j++ {
			new_mat[i][j] = matrix[j][i]
		}
	}
	return new_mat
}
