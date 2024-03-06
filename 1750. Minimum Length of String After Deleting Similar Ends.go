// Time C. O(n), Space C. O(1)
// runtime 


func minimumLength(s string) (res int) {
	l, r := 0, len(s)-1

	for l < r && s[l] == s[r] {
		char := s[l]
		for l <= r && s[l] == char {
			l++
		}
		for r >= l && s[r] == char {
			r--
		}
	}

	return r - l + 1
}

