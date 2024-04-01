// started a bit of js
// Time C. O(n), Space C. O(1)
// Runtime 1ms

import "strings"

func lengthOfLastWord(s string) int {
	out := strings.TrimRight(s, " ")
	conn := strings.Split(out, " ")
	return len(conn[len(conn) - 1])
}