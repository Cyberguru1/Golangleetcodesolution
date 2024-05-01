// Time C. O(n), Space C. O(1)
// Runtime 1ms

func reversePrefix(word string, ch byte) (res string) {
	res = word

	rev := func(word string) (res string) {
		for _, k := range word {
			res = string(k) + res
		}
		return
	}

	for i, d := range word {
		if ch == byte(d) {
			res = rev(word[:i+1]) + word[i+1:]
			return
		}
	}
	return
}
