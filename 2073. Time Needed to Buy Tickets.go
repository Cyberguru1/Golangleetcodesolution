// Time C. O(n^2), Space C. O(1)
// BestRuntime 2ms
func timeRequiredToBuy(tickets []int, k int) (count int) {

	for tickets[k] > 0 {
		for i, _ := range tickets {
			if tickets[i] > 0 {
				tickets[i]--
				count++
			}
			if i == k && tickets[i] == 0 {
				break
			}
		}
	}

	return count
}
