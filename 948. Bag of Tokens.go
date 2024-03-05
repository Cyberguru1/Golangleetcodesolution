// Time C. O(n), Space C. O(n)
// runtime 5ms

func bagOfTokensScore(tokens []int, power int) (res int) {

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})

	res = 0

	l := 0
	h := len(tokens) - 1

	for l <= h {
		if tokens[l] <= power {
			res++
			power -= tokens[l]
			l++
		} else if (tokens[h] > power) && (res > 0) && (l != h) {
			res--
			power += tokens[h]
			h--
		} else {
			return
		}
	}

	return res
}
