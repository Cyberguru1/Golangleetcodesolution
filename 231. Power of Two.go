// Time C. O(1), Space C. O(1)
// runtime 

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	res := n & (n - 1)

	if res == 0 {
		return true
	}
	return false
}