// Time C. O(n) Space C. O(N)
// Best runtime 0ms

// man be writing too much for loop this days sgn

func customSortString(order string, s string) (res string) {
	mapp := map[rune]int{}

	for _, k := range s {
		mapp[k]++
	}

	for _, k := range order {
		if _, ok := mapp[k]; ok {
			res += string(k)
			mapp[k]--
		}
		for i, _ := mapp[k]; i > 0; i-- {
			res += string(k)
			mapp[k]--
		}
	}

	for v, c := range mapp {

		for i := c; i > 0; i-- {
			res += string(v)
		}
	}

	return
}