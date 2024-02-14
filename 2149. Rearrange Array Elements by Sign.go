// Time C. O(n), Space C. O(n^2)
// runtime 222ms

func rearrangeArray(nums []int) (res []int) {
	arr1 := []int{}
	arr2 := []int{}

	for _, k := range nums {
		if k > 0 {
			arr1 = append(arr1, k)
		} else {
			arr2 = append(arr2, k)
		}
	}

	for i := 0; i < len(arr1); i++ {
		res = append(res, arr1[i], arr2[i])
	}

	return
}

// second solution using only ans array
// Time C. O(n), Space C. O(n)
// runtime 184ms

func rearrangeArray(nums []int) (res []int) {

	res = make([]int, len(nums))

	for i, j, k := 0, 1, 0; k < len(nums); k++ {

		if nums[k] > 0 {
			res[i] = nums[k]
			i += 2
		} else {
			res[j] = nums[k]
			j += 2
		}

	}

	return
}

// took 5mins to solve both soln, improving !!!