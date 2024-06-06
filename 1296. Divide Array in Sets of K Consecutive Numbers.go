// Time C. O(n^2), Space C. O(n)
// Runtime 1333ms

func isPossibleDivide(nums []int, k int) bool {
	lent := len(nums)

	if (lent / k) * k != lent {
		return false
	}

	sort.Slice(nums, func(i, j int)bool{
		return nums[i] < nums[j]
	})

	hashmap := map[int]int{}

	for _, v := range nums {
		hashmap[v]++
	}

	for len(nums) > 0 {
		v := nums[0]
		for ii := 0; ii < k; ii++ {
			if vv, ok := hashmap[v+ii]; !ok || vv <= 0 {
				return false
			}else {
				hashmap[v+ii]--
				for l:= 0; l < len(nums); l++ {
					if nums[l] == v+ii {
						nums = append(nums[:l], nums[l+1:]...)
						break
					}
				}
			}
		}
	}

	return true

}