// Time C. O(n*k) where n is length of words and k is len of each word in words
// Space C. O(1) since input is within the range of 26 alphabets
// Best runtime 5ms

func makeEqual(words []string) bool {

	if len(words) == 1 {
		return true
	}

	chars := make([]int, 26)

	var makeCnt func(s string)

	makeCnt = func(s string) {
		for i := 0; i < len(s); i++ {
			chars[s[i]-'a']++
		}
	}

	for _, d := range words {
		makeCnt(d)
	}

	for _, v := range chars {
		if v > 0 {
			if v%len(words) != 0 {
				return false
			}
		}
	}

	return true
}