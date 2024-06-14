// Time C. O(NlogN), Space O(n)

func minIncrementForUnique(nums []int)(count int ){
	sort.Ints(nums)
	n := len(nums)
	maxV := nums[len(nums)-1]
	freq := make([]int, (n+maxV+1))

	for _, v := range nums {
		freq[v]++
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] <= 1 {
			continue
		}

		dup := freq[i] - 1
		freq[i+1] += dup
		freq[i] = 1
		count += dup
	}
	return
}
