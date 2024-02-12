// Time C. O(n), Space C. O(n)
// runtime 12ms

func majorityElement(nums []int) int {

	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	mapp := map[int]int{}

	for _, k := range nums {
		mapp[k]++
		if kk, ok := mapp[k]; ok {
			if kk == (n/2)+1 {
				return k
			}
		}
	}

	return 0
}

// Follow up using O(1) space complexity

// Time C. O(nlogn), Space C. O(1)
// runtime 11ms

func majorityElement(nums []int) int {

	n := len(nums)

	if n == 1 {
		return nums[0]
	}
	n = (n / 2) + 1

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	c := 1
	init := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == init {
			c++
		} else {
			init = nums[i]
			c = 1
		}

		if c == n {
			return nums[i]
		}
	}
	return 0
}