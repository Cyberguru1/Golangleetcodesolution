// Time C. O(log(n)), Space C, O(1)
// runtime 32ms

func reverseString(s []byte) {
	// var p byte
	lenn := len(s)
	for i := 0; i < lenn/2; i++ {
		// p = s[i]
		// s[i] = s[lenn - i - 1]
		// s[lenn - i - 1] = p

		s[i], s[lenn-i-1] = s[lenn-i-1], s[i]
	}
}