// Time C. O(n * 2) Space C. O(n)
// runtime 40ms907. Sum of Subarray Minimums


func sumSubarrayMins(arr []int) int {
	stack, res, mod := []int{}, 0, 1000000007
	for j := 0; j < len(arr); j++ {
		for len(stack) > 0 && arr[j] < arr[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			right := j
			count := (mid - left) * (right - mid) % mod
			res += (arr[mid] * count) % mod
			res %= mod
		}
		stack = append(stack, j)
	}
	right := len(arr)
	for len(stack) > 0 {
		mid := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := -1
		if len(stack) > 0 {
			left = stack[len(stack)-1]
		}
		count := (mid - left) * (right - mid) % mod
		res += (arr[mid] * count) % mod
		res %= mod
	}
	return res
}