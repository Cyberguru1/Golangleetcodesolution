// Time C. O(n), Space C. O(1)
// Runtime 1ms

func checkValidString(s string) bool {
	lmin, lmax := 0, 0

	for _, v := range s {
		if v == '(' {
			lmin++
			lmax++
		} else if v == ')' {
			lmin--
			lmax--
		} else {
			lmin--
			lmax++
		}

		if lmax < 0 {
			return false
		} else if lmin < 0 {
			lmin = 0
		}
	}
	if lmin == 0 {
		return true
	}
	return false
}
