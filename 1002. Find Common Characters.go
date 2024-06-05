// Time C. O(NM), Space C. O(nm), where n is the len of words and m is the len of word in words
// Runtime 5ms

func commonChars(words []string) (res []string) {

	hashmap := map[int]map[rune]int{}

	for i := 0; i < len(words); i++ {
		hashmap[i] = map[rune]int{}
		for _, k := range words[i] {
			hashmap[i][k]++
		}
	}

	for k, v := range hashmap[0] {
		count, minn := 1, v
		for i := 1; i < len(words); i++ {
			if vv, ok := hashmap[i][k]; ok {
				count++
				minn = min(minn, vv)
			}
		}
		if count == len(words) {
			for i := 0; i < minn; i++ {
				res = append(res, string(k))
			}
		}
	}

	return
}
