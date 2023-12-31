// Time C. O(n), Space C. O(1)
// best runtime 0ms

func maxLengthBetweenEqualCharacters(s string) int {
	maxx := -1
	mapp := make(map[rune]int)

	for i, v := range s {
		if _, ok := mapp[v]; ok {
			maxx = max(maxx, i-mapp[v])
		} else {
			mapp[v] = i + 1
		}
	}

	return maxx
}
