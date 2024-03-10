// first approach
// Time C. O(nlogn), Space C. O(1)
// Best runtime 4ms

func intersection(nums1 []int, nums2 []int) (res []int) {

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	l, j := 0, 0

	for l < len(nums1) && j < len(nums2) {
		if nums1[l] == nums2[j] {
			if len(res) > 0 && res[len(res)-1] == nums1[l] {
				l++
				j++
				continue
			} else {
				res = append(res, nums1[l])
			}
		}
		if nums1[l] < nums2[j] {
			l++
		} else {
			j++
		}
	}
	return

}

// second apprach
// Time c. O(n), Space C. O(n)
// Best runtime 0ms

func intersection(nums1 []int, nums2 []int) (res []int) {

	mapps := map[int]bool{}

	for _, k := range nums1 {
		mapps[k] = true
	}

	for _, j := range nums2 {
		if out, ok := mapps[j]; ok && out {
			res = append(res, j)
			mapps[j] = false
		}
	}

	return

}