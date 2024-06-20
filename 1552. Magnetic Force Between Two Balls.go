// Time C O(n), Space C. O(1)

func maxDistance(position []int, m int) (res int ){
	sort.Ints(position)

	check := func(x int) bool {
		prev_pos := position[0]
		place := 1

		for _, val := range position[1:] {
			curr_pos := val
			if curr_pos - prev_pos >= x {
				place++
				prev_pos = curr_pos
			}

			if place == m {
				return true
			}
		}
		return false
	}

	lent := len(position)
	l, h := 1, position[lent-1] / (m - 1)
	h++
	for l <= h {
		mid := l + (h - l)/2

		if check(mid) {
			res = mid
			l = mid + 1
		}else {
			h = mid - 1
		}
	}

	return

}