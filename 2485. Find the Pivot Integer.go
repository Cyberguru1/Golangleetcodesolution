// Runtime 0ms
// Time C. O(logn), space C. O(1)

func pivotInteger(n int) ( res int ){
	sum := n * (n + 1) / 2

	l, r := 1, n
	for l <= r {
		m := (l + r) / 2
		firstPart := m * (m + 1) / 2
		secondPart := sum - firstPart + m

		if firstPart == secondPart {
			return m
		} else if firstPart > secondPart {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}