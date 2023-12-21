// first brute force solution,
// Time C. O(nlogn), Space C. O(1)
// best runtime 116ms, Space 14Mb, Beats 100% of User in Go



func maxWidthOfVerticalArea(points [][]int) int {

	slices.SortFunc(points,
					func(Fpoint, Spoint []int) int {
						return cmp.Compare(Fpoint[0], Spoint[0])
					})

	Amax := 0

	for i:= 0; i < (len(points) - 1); i++ {
		Amax = max((points[i+1][0] - points[i][0]), Amax)
	}
	return Amax
}
