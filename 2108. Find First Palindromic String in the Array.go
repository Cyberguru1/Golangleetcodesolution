// Time C. O(n), Space C. O(1)
// runtime 15ms

func firstPalindrome(words []string) string {

	check := func(str string) bool {

		for i := 0; i < len(str)/2; i++ {
			if str[i] != str[len(str)-i-1] {
				return false
			}
		}
		return true
	}

	for _, k := range words {
		if check(k) {
			return k
		}
	}

	return ""

}