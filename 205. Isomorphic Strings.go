// Time C. O(n), Space C. O(n)
// Best runtim 0ms

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	mapp := map[byte]byte{}
	mapp1 := map[byte]byte{}

	for i := 0; i < len(s); i++ {

		if d, ok := mapp[s[i]]; ok && d != t[i] {
			return false
		}else{
			mapp[s[i]] = t[i]
		}

		if d, ok := mapp1[t[i]]; ok && d != s[i] {
			return false
		}else{
			mapp1[t[i]] = s[i]
		}
	}
	return true    
}
