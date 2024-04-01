// runtime 96 ms
// Time C. O(N), Space C. O(1)

func countSubarrays(nums []int, minK int, maxK int) (ans int64) {
	lastMin, lastMax := -1, -1
	leftmost := 0
	for i, d := range nums {
		if d < minK || d > maxK {
			leftmost = i + 1
			lastMin, lastMax = -1, -1
			continue
		}
		if d == minK {
			lastMin = i
		}
		if d == maxK {
			lastMax = i
		}
		if lastMin != -1 && lastMax != -1 {
			ans += int64(min(lastMin, lastMax)-leftmost+1)
		}
	}

	return 
}