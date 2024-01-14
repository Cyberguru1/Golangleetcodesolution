// Time C. O(n), Space C. O(1)
// runtime 74ms

func closeStrings(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	dic := map[byte]int{}

	for k, _ := range word1 {
		dic[word1[k]]++
	}

	dic1, dic2 := make([]byte, 26), make([]byte, 26)

	for k, _ := range word1 {
		dic1[word1[k]-'a']++
		dic2[word2[k]-'a']++

		if _, ok := dic[word2[k]]; !ok {
			return false
		}
	}

	slices.Sort(dic1)
	slices.Sort(dic2)

	if slices.Equal(dic1, dic2) {
		return true
	}

	return false

}