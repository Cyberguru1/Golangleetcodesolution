// Time C. O(n), Space C. O(1)
// Runtime 99ms

func averageWaitingTime(customers [][]int) (res float64) {
	if len(customers) == 0 {
		return res
	}
	fin := customers[0][0] + customers[0][1]
	count, total := 1, fin-customers[0][0]
	for _, customer := range customers[1:] {
		if customer[0] > fin {
			fin = customer[0] + customer[1]
		} else {
			fin = fin + customer[1]
		}
		total += (fin - customer[0])
		count++
	}
	res = float64(total) / float64(count)
	return

}