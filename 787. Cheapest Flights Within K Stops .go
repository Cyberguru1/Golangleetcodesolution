// welcome to more neetcode :laughs:
// Time C. O(E.k) Space C. O(E.k) where e is number of edge and k is number of stops
// runtime 7ms

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]float64, n)
	for i := 0; i < n; i++ {
		prices[i] = math.Inf(0)

	}

	prices[src] = 0

	for i := 0; i < k+1; i++ {
		temp := slices.Clone(prices)

		for _, d := range flights {
			s, d, p := d[0], d[1], d[2]

			if prices[s] == math.Inf(0) {
				continue
			}

			if prices[s]+float64(p) < temp[d] {
				temp[d] = prices[s] + float64(p)
			}
		}

		prices = temp
	}

	if prices[dst] == math.Inf(0) {
		return -1
	}
	return int(prices[dst])
}