// brute force solution 4, time C. : O(n), space C. :O(26)
// best runtime 4ms

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	chars := make(map[byte]int, 26)

	for i, _ := range s {
		chars[s[i]]++
		chars[t[i]]--
	}

	for _, v := range chars {
		if (v != 0){
			return false
		}
	}

	return true

}