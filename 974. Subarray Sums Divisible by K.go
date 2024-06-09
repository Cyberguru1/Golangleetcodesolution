// Time C. O(n), Space C. O(n)
// Runtime 4ms

func subarraysDivByK(nums []int, k int) int {
	prefix_map := make(map[int]int)
	prefix_map[0] = 1
	sum, ans := 0, 0
	for i := 0; i < len(nums); i++ {
		sum = (sum + nums[i]) % k
		if sum < 0 {
			sum += k
		}
		if _, ok := prefix_map[sum]; ok {
			ans += prefix_map[sum]
			prefix_map[sum]++
		} else {
			prefix_map[sum] = 1
		}
	}
	return ans
}
