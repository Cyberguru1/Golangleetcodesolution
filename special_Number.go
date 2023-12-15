// First bruteforce solution (7ms)

func numSpecial(mat [][]int) int {
	result := 0
	lrow := len(mat)
	lcol := len(mat[0])

   var col_arr []int
   var row_arr []int

	for i:=0; i < lrow; i++ {
		rsum := 0
		for j:= 0; j < lcol; j++ {
			rsum += mat[i][j]
		}
		col_arr = append(col_arr, rsum)
	}

	for i:=0; i < lcol; i++ {
		csum := 0
		for j:= 0; j < lrow; j++ {
			csum += mat[j][i]
		}
		row_arr = append(row_arr, csum)
	}


	for i,v := range mat {
		if (col_arr[i] != 1){
			continue
		} else {
			for j, k := range v {
				if (k == 1) && (row_arr[j] == k) {
					result += 1
				}
			}
		}
	}
	return result
}