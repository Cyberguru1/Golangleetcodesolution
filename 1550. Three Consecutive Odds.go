

func threeConsecutiveOdds(arr []int) (res bool) {

	count := 0

	for _, val := range arr {

		if val%2 == 1 {
			count++
		} else {
			count = 0
		}

		if count == 3 {
			res = true
			return
		}
	}

	return
}