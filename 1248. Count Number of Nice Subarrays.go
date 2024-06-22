// Time C. O(n), Space C. O(n)

func numberOfSubarrays(nums []int, k int) (res int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			count++
			nums[i] = 0
		} else {
			nums[i] = 1
		}
	}

	if len(nums) == count {
		return
	}

	hashmap := map[int]int{0: 1}
	pSum := 0

	for i := 0; i < len(nums); i++ {

		pSum += nums[i]
		diff := pSum - k

		if v, Ok := hashmap[diff]; Ok {
			res += v
		}

		hashmap[pSum]++
	}

	return
}