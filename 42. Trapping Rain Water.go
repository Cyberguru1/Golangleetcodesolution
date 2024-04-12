// a tuff run
// Time C. O(n), Space C. O(N)
// Runtime 11ms

func trap(height []int)(res int ){
	lent := len(height)

	if lent == 0 { return 0}

	l, r := make([]int, lent), make([]int, lent)


	l[0], r[lent - 1] = height[0], height[lent - 1]

	for i:= 1; i < lent; i++{
		l[i] = max(l[i-1], height[i])
	}

	for i := lent - 2; i > -1; i--{
		r[i] = max(r[i+1], height[i])
	}

	for i, _ := range height {
		res += min(r[i], l[i]) - height[i]
	}

	return 

}