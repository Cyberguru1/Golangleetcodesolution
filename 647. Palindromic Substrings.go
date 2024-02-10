// Time C. O(n*2), Space C. O(1)
// runtime 58 ms

func countSubstrings(s string) (res int) {
	res = len(s)

	if res == 1 {
		return
	}

	checkp := func(s string) bool {

		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}

	// use sliding window with incremental i as window
	for i := 2; i < len(s); i++ {
		for j := 0; j < len(s)-(i-1); j++ {
			if checkp(s[j : j+i]) {
				res++
			}
		}
	}

	if checkp(s) {
		res++
	}

	return
}