// Time C. O(n), Space C. O(1)
// Runtime 0ms

func scoreOfString(s string) (res int) {
	abs := func(val int) int {
		return int(math.Abs(float64(val)))
	}

	for i := 0; i < len(s)-1; i++ {
		res += abs(int(rune(s[i]) - rune(s[i+1])))
	}
	return
}