// Time C. O(2^n), Space C. O(n)

func wordBreak(s string, wordDict []string)(result []string) {

	mapp := make(map[string]int, len(wordDict))

	for _, k := range wordDict {
		mapp[k] = 0
	}

	valid := func(word []string) bool {
		res := ""
		res1 := ""

		for i := 0; i < len(word); i++ {
			if i != len(word) - 1 {
				res += word[i]
				res1 += word[i] + " "
			}else{
				res += word[i]
				res1 += word[i]
			}
		}

		if res == s {
			result = append(result, res1)
			return true
		}

		return false
	}

	var backtrack func(currWords []string, st int)

	backtrack = func(currWords []string, st int) {

		if len(currWords) != 0 && valid(currWords) {
			return
		}
		for i := st; i < len(s); i++ {
			res := ""
			for j := i; j < len(s); j++ {
				res += string(s[j])
				if _, ok := mapp[res]; ok {
					currWords = append(currWords, res)
					backtrack(currWords, j+1)
					currWords = currWords[:len(currWords)-1]
				}
			}
		}
	}

	backtrack([]string{}, 0)

	return
}