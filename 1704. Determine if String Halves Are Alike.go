// Time C. O(n), Space C. O(1)
// runtime 0ms

type void struct{}

func halvesAreAlike(s string) bool {

	mVowels := map[byte]void{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}

	count := 0

	for i, _ := range s {
		if i < (len(s) / 2) {
			if _, ok := mVowels[s[i]]; ok {
				count++
			}
		} else {
			if _, ok := mVowels[s[i]]; ok {
				count--
			}

		}

	}

	if count != 0 {
		return false
	}

	return true

}