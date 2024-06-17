// Time C. O(n), Space C. O(1)

func judgeSquareSum(c int) bool {

	l, r := 0, int(math.Pow(float64(c), 0.5))
	var val int64

	for (l * l) <= c && (r * r) >= 0 {

	   val  = int64((l * l) + (r * r))

		if  val == int64(c) {
			return true
		}

		if val > int64(c) {
			r--
			continue

		}

		l++
	}

	return false

}