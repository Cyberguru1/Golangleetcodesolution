
// first brute force solution
// time C. O(n*m), space C. O(n*m)
// Ugly if conditions <(-_-)>

func imageSmoother(img [][]int) [][]int {
	lr, lc := len(img), len(img[0])

	new_mat := make([][]int, lr)
	var sum, cnt int

	for i, v := range img {
		new_mat[i] = make([]int, lc)
		for j, _ := range v {
			sum = img[i][j]
			cnt = 1

			if ((i-1) >= 0){
				sum += img[i-1][j]
				cnt++

				if ((j-1) >= 0){
					sum += img[i-1][j-1]
					cnt++
				}
				if ((j+1) < lc){
					sum += img[i-1][j+1]
					cnt++
				}
			}
			if ((j - 1) >= 0){
				sum += img[i][j-1]
				cnt++

				if ((i + 1) < lr){
					sum += img[i+1][j-1]
					cnt++
				}
			}
			if ((i + 1) < lr){
				sum += img[i+1][j]
				cnt++
			}
			if ((j+1) < lc){
				sum += img[i][j+1]
				cnt++
				if ((i+1) < lr){
					sum += img[i+1][j+1]
					cnt++
				}
			}

			new_mat[i][j] = (sum/cnt)
		}

	}
	return new_mat
}