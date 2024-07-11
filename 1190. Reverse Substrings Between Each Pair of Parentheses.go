// Time C. O(n^2), Space C. O(n)

func reverseParentheses(s string) string {

	o, c := '(', ')'
	stack := ""

	for _, k := range s {
		if k == c {
			con := ""
			i := len(stack) - 1
			for stack[i] != byte(o) {
				con += string(stack[i])
				i--
			}
			stack = stack[:i] + con
			continue
		}
		stack += string(k)
	}
	return stack
}
