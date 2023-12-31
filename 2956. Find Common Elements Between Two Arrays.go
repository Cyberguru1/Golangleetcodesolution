
// first brute force solution
// Time C. O(n^2)
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	res := make([]int, 2)

	for _, j := range nums1 {
		check := make(map[int]bool)
		for _, k := range nums2 {
			if (j == k) && check[j] != true {
				res[0]++
				check[j] = true
			}
		}
	}

	for _, j := range nums2 {
		check := make(map[int]bool)
		for _, k := range nums1 {
			if (j == k) && check[j] != true {
				res[1]++
				check[j] = true
			}
		}
	}

	return res
}

// second brute force solution
// Time C. O(n)

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	res := make([]int, 2)
	dic1 := make(map[int]int)
	dic2 := make(map[int]int)

	for _, v := range nums1 {
		dic1[v]++
	}

	for _, v := range nums2 {
		dic2[v]++
	}

	for v, _ := range dic1 {
		if _, ok := dic2[v]; ok {
			res[1] += dic2[v]
			res[0] += dic1[v]
		}
	}

	return res
}