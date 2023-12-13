package main

import (
	"fmt"
)

func climbStairs(n int, store map[int]int) int {
	
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}

	check := true
	for key, _ := range store {
		if n == key {
			check = false
			return store[key]
		}
	}
	if check {
		store[n] = climbStairs(n-1, store) + climbStairs(n-2, store)
	}
	return climbStairs(n-1, store) + climbStairs(n-2, store)

}



func main() {
	x := map[int]int{}
	fmt.Println(climbStairs(9000, x))
}
