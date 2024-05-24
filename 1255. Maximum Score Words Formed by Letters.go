// Time C. O(2^n), Space C. O(n)


func maxScoreWords(words []string, letters []byte, score []int) (maxx int) {

	mapp := make(map[byte]int, len(letters))

	for _, k := range letters {
		mapp[k]++
	}

	add := func(word string) {
		for _, j := range word {
			mapp[byte(j)]++
		}
	}

	sub := func(word string) {
		for _, j := range word {
			mapp[byte(j)]--
		}
	}

	valid := func(word string) bool {
		for _, k := range word {
			if v, _ := mapp[byte(k)]; v < 0 {
				return false
			}
		}
		return true
	}

	calc := func(word string) (summ int) {
		for _, j := range word {
			summ += score[(rune(j) - 97)]
		}
		return
	}

	var backtrack func(count, ind int)

	backtrack = func(count, ind int) {

		maxx = max(count, maxx)

		for j := ind; j < len(words); j++ {
			val := calc(words[j])
			sub(words[j])
			if valid(words[j]) {
				count += val
				backtrack(count, j+1)
				count -= val
			}
			add(words[j])

		}
	}

	backtrack(0, 0)
	return
}