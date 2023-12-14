// Bruteforce solution (7ms):

func maxProduct(nums []int) int {

	var result int = 0

	for jj, j := range nums {
		for ii, i := range nums {
			val := (i - 1)*(j - 1)
			if (val > result) && (jj != ii){
				result = val
			}
		}
	}
	return result
}

// Second bruteforce solution (6ms):

func maxProduct(nums []int) int {

	var result int = 0

	if len(nums) == 2 {
		return (nums[0] - 1) * (nums[1] - 1)
	}
	for i:= 1; i < len(nums); i++ {
		for j:= 0; j< i ; j++ {
			val := (nums[i] - 1)* (nums[j] - 1)
			if val > result {
				result = val
			}
		}
	}
	return result
}


// Third brute force solution (4ms):

func maxProduct(nums []int) int {

	var result int = 0

	if len(nums) == 2 {
		return (nums[0] - 1) * (nums[1] - 1)
	}
	max, maxi := 0, 0

	for i, j := range nums {
		if j > max {
			max, maxi = j, i
		}
	}

	for i, j := range nums {
		val := (max - 1) * (j - 1)
		if (i != maxi) && (val > result){
			result = val
		}

	}
	return result
}

// simillar to the previous (shorter approach) (4ms):

func maxProduct(nums []int) int {
	max1, max2 := 0, 0

	for _, j := range nums{
		if (j > max1){
			max2 = max1
			max1 = j
		} else if (j > max2){
			max2 = j
		}
	}

	return (max1 - 1) * (max2 -1 )
}