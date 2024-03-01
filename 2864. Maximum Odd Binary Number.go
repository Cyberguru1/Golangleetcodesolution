// intuitive as possible
// runtime O(n), Space O(n)

func maximumOddBinaryNumber(s string) string {
	return strings.Repeat("1", strings.Count(s, "1")-1) + strings.Repeat("0", strings.Count(s, "0")) + "1"
}
