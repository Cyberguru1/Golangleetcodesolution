// Time C. O(n), Space C. O(n)
// runtime 5ms

import  "strconv"
func evalRPN(tokens []string) int {

	lstack := []string{}

	ops := map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
	}

	for _, k := range tokens{
		if _, ok := ops[string(k)]; !ok{
			lstack = append(lstack, string(k))
		}else{
			v1,_ := strconv.Atoi(lstack[len(lstack)-1])
			lstack = lstack[:len(lstack)-1]
			v2, _ := strconv.Atoi(lstack[len(lstack)-1])
			lstack = lstack[:len(lstack)-1]
			lstack = append(lstack, fmt.Sprint(ops[string(k)](v2, v1)))
		}
	}
	out, _ := strconv.Atoi(lstack[0])
	return   out
}