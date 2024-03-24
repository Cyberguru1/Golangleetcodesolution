// TIme C. O(nlogn), Space C. O(1)
// runtime 117ms
// yeah a lazy soln

func findDuplicate(nums []int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return 0

}

// TO-DO, since array nums contains val within [1, n] and len(n+1)
// then indexing the array nums[nums[i]] is also valid within the array
//

