// Time C. O(n), Space C. O(n)
// Runtime 119ms

func checkSubarraySum(nums []int, k int) bool {
	sum := nums[0]
	hashmap := map[int]bool{}
	hashmap[nums[0]] = true

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if (nums[i] + nums[i-1])%k == 0  {
			return true
		}else if _, ok := hashmap[sum%k]; ok && (sum-nums[i]) % k != sum%k || (sum%k) == 0 {
			return true
		}
		hashmap[sum%k] = true
	}

	return false
}