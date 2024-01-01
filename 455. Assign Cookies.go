// Time C. O(nlogn) <- due to sorting, Space C. O(m+n)

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	ls, lg := 0, 0

	for (ls < len(s)) && (lg < len(g)) {
		if s[ls] >= g[lg] {
			lg++
		}
		ls++
	}

	return lg
}455. Assign Cookies
