// Time C. O(n), Space C. O(1)
// best runtime 7ms

func minSteps(s string, t string) int {
	smap := make([]int, 26)

	for v, _ := range s {
		smap[s[v]-'a']++
		smap[t[v]-'a']--
	}

	count := 0

	for _, k := range smap {
		if k > 0 {
			count += k
		}
	}

	return count
}