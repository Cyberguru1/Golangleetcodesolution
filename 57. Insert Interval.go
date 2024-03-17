// runtime 8ms
// Time C. O(n), same with space

func insert(intervals [][]int, newInterval []int) [][]int {

	output := [][]int{}

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			output = append(output, newInterval)
			return append(output, intervals[i:]...)
		} else if newInterval[0] > intervals[i][1] {
			output = append(output, intervals[i])
		} else {
			newInterval = []int{
				min(newInterval[0], intervals[i][0]),
				max(newInterval[1], intervals[i][1]),
			}
		}
	}
	output = append(output, newInterval)
	return output
}