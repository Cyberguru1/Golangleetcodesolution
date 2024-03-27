// Time C. O(n^2), Space C. O(n)
// runtime 63ms

func numSubarrayProductLessThanK(nums []int, k int) (count int) {

	lent := len(nums)

	for i := 0; i < lent; i++ {
		prod := 1
		for j := i; j < lent; j++ {
			prod *= nums[j]
			if prod < k {
				count++
			} else if prod > k {
				break
			}
		}
	}
	return
}
