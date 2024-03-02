// Time C. O(nlogn), Space C. O(n)
// runtime 30ms
func sortedSquares(nums []int)( res []int) {

	for _, k := range nums{
		res = append(res, k*k)        
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return
}

// Follow up using O(n) solution
// Time C. O(n), Space C. O(n)

func sortedSquares(nums []int) []int {

	l := 0
	h := len(nums)-1
	pos := h

	res := make([]int, h+1)

	for l <= h {
		if nums[h] > int(math.Abs(float64(nums[l]))){
			res[pos] = nums[h] * nums[h]
			pos--
			h--
		}else {
			val := int(math.Abs(float64(nums[l])))
			res[pos] = val * val
			pos--
			l++
		}
	}

	return res
}