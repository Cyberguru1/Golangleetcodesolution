// Time C O(n), Space C. O(1)

func intersect(nums1 []int, nums2 []int) (res []int) {

	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := len(nums1)-1, len(nums2)-1

	for i >= 0 && j >= 0 {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i--
			j--
			continue
		}

		if nums1[i] > nums2[j]{
			i--
			continue
		}

		j--
	}
	return
}