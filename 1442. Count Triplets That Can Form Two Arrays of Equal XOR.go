// Time C. O(n^2), Space C. O(n)

func countTriplets(arr []int) int {
	n := len(arr)
	count := 0
	prefixXOR := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefixXOR[i+1] = prefixXOR[i] ^ arr[i]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				if (prefixXOR[j]^prefixXOR[i]) == (prefixXOR[k+1]^prefixXOR[j]) {
					count++
				}
			}
		}
	}

	return count
}