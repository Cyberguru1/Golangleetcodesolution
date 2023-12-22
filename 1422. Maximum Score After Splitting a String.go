// first brute force solution
// Time C. O(n), Space C. O(1)
// runtime 0ms beats 100%, space 2Mb

func maxScore(s string) int {

	C_ones := strings.Count(s, "1")

	C_zero, Rmax := 0, 0


	for i, v := range s {
		if (v == '0') && (i != len(s) - 1){
			C_zero++
			csum := C_zero + C_ones
			Rmax = max(Rmax, csum)
		} else {
			C_ones--
			csum := C_zero + C_ones
			Rmax = max(Rmax, csum)
		}
	}

	return Rmax
}