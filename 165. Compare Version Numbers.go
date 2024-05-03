// Time C. O(N), Space C. O(1)
// Runtime 0ms

func compareVersion(version1 string, version2 string) int {

	nums1 := strings.Split(version1, ".")
	nums2 := strings.Split(version2, ".")

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2){

		k1, _ := strconv.Atoi(nums1[i])
		k2, _ := strconv.Atoi(nums2[j])

		if k1 < k2 {
			return -1
		} else if k1 > k2 {
			return 1
		}
		i++
		j++
	}

	for i < len(nums1) {
		k1, _ := strconv.Atoi(nums1[i])
		if k1 > 0 {
			return 1
		}
		i++
	}
	for j < len(nums2) {
		k2, _ := strconv.Atoi(nums2[j])
		if k2 > 0 {
			return -1
		}
		j++
	}
	return 0

}