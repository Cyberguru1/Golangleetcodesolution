// Time C. O(n), Space C. O(n)
// runtime 86ms

type void struct {
}

func findJudge(n int, trust [][]int) int {

	if len(trust) == 0 && n == 1 {
		return n
	}
	cont := map[int]void{}
	tcont := map[int]int{}

	var d void

	for _, k := range trust {
		cont[k[0]] = d
		tcont[k[1]]++
	}

	for _, k := range trust {
		if _, ok := cont[k[1]]; !ok && tcont[k[1]] == n-1 {
			return k[1]
		}
	}

	return -1
}