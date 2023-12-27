// runtime 111ms
// Time C. O(n), Space C. O(1)

func minCost(colors string, neededTime []int) int {
	l := colors[0]
	lt := neededTime[0]
	min_t := 0

	for i := 1; i < len(colors); i++ {
		if ((l == colors[i]) && (lt >= neededTime[i])){
			min_t += neededTime[i]
		} else if ((l == colors[i]) && (lt < neededTime[i])){
			min_t += lt
			lt = neededTime[i]
		} else{
			lt = neededTime[i]
			l = colors[i]
		}
	}
	return min_t
}