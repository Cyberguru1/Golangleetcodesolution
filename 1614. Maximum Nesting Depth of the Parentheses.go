// Time C. O(n), Space C. O(1)
// Runtime 1ms

func maxDepth(s string) (count int) {
	cun := 0
	for _, d := range s {
		if string(d) == "(" {
			cun++
			count = max(cun, count)
		} else if string(d) == ")" {
			cun--
		}
	}
	return
}