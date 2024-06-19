// Time O(nk), Space C. O(n)

func minDays(bloomDay []int, m int, k int) (res int) {

	if int64(len(bloomDay)) < int64(m*k) {
		return -1
	}

	lent := len(bloomDay)

	sortBloomDay := append([]int{}, bloomDay...)
	sort.Ints(sortBloomDay)

	st := 0
	prev := 0
	for st < lent {
		currDay := sortBloomDay[st]
		if currDay == prev {
			st++
			continue
		}

		count := 0
		nb := 0

		for i := 0; i < lent; i++ {
			if currDay >= bloomDay[i] {
				count++
			} else {
				count = 0
				continue
			}
			if count == k {
				nb++
				count = 0
			}
			if int64(nb) >= int64(m) {
				res = currDay
				break
			}
		}
		if res == currDay {
			break
		}
		prev = currDay
		st++
	}
	return
}