// Time C. O(nlogn), Space C. O(1)
// Runtime 2ms

func numRescueBoats(people []int, limit int) int {

	slices.Sort(people)
	count := 0
	i, l := 0, len(people)-1

	for i <= l {

		if people[i]+people[l] <= limit {
			count++
			i++
		} else {
			count++
		}
		l--
	}

	return count
}