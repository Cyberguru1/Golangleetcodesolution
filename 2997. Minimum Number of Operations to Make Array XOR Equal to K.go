// Time C. O(1), Space C O(1)
// runtime 10ms
func minOperations(nums []int, k int) int {
	res := k
	for _, n := range nums {
		res ^= n
	}
	var ans int
	for res > 0 {
		ans += res % 2
		res = res >> 1
	}
	return ans
}

