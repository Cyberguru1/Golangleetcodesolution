//

func climbStairs(n int) int {
	x := map[int]int{}
	return ClimbStairs(n, x)

}

func ClimbStairs(n int, store map[int]int) int {

	if n <= 3 {
		return n
	}

	check := true
	for key, _ := range store {
		if n == key {
			check = false
			return store[key]
		}
	}
	if check {
		store[n] = ClimbStairs(n-1, store) + ClimbStairs(n-2, store)
	}
	return ClimbStairs(n-1, store) + ClimbStairs(n-2, store)
}
