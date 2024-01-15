// Time C. O(n), Space C. O(n)

func findWinners(matches [][]int) [][]int {

	wmapp := map[int]int{}
	lmapp := map[int]int{}

	for _, k := range matches {
		wmapp[k[0]]++
		lmapp[k[1]]++
	}

	ans := make([][]int, 2)
	ans[0] = []int{}
	ans[1] = []int{}

	for k, _ := range wmapp {
		if _, ok := lmapp[k]; !ok {
			ans[0] = append(ans[0], k)
		}
	}

	for k, _ := range lmapp {
		if lmapp[k] == 1 {
			ans[1] = append(ans[1], k)
		}
	}
	slices.Sort(ans[0])
	slices.Sort(ans[1])
	return ans
}