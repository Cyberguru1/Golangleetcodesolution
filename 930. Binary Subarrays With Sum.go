// doing pico, didn't have time to implement the best algo, but this works <-_->
// Time C. O(n^2) Space C. O(1)

func numSubarraysWithSum(nums []int, goal int) (res int) {

	for i := 0; i < len(nums); i++ {
		sum := 0
		for k := i; k < len(nums); k++ {
			sum += nums[k]
			if sum == goal {
				res++
			} else if sum > goal {
				break
			}
		}
	}

	return
}
