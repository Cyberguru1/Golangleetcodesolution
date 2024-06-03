// Time C. O(n), Space C. O(1)
// Runtime 3ms

func appendCharacters(s string, t string) int {

	count, l, r := 0, 0, 0

	for l < len(s) && r < len(t){
		if s[l] == t[r] {
			count++
			l++
			r++
			continue
		}
		l++
	}

	return (len(t) - count)   
}
