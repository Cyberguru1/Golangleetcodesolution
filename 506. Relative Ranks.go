// Time C. O(nlogn) Space C. O(n)
// Runtime 15ms

func findRelativeRanks(score []int)(res []string ){

	if len(score) <= 2 {

		if len(score) == 1 {
			return []string{"Gold Medal"}
		}
		if score[1] > score[0]{
			return []string{"Silver Medal", "Gold Medal"}
		}

		return []string{"Gold Medal", "Silver Medal"}
	}

	init := make([]int, len(score))
	copy(init, score)

	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	mapps := map[int]string{score[0]:"Gold Medal", score[2]:"Bronze Medal", score[1]:"Silver Medal"}

	k := 4
	for _, v := range score[3:] {
		mapps[v] = strconv.Itoa(k)
		k++
	}

	for _, v  := range init {
		out, _ := mapps[v]
		res = append(res, out)
	}

	return
}