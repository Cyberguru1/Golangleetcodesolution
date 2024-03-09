// Time C. O(n+m), Space C. O(1)
// best runtime 53ms

func getCommon(nums1 []int, nums2 []int) int {

	p1 := 0
	p2 := 0

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			return nums1[p1]
		}else if nums1[p1] < nums2[p2]{
			p1++
		}else{
			p2++
		}
	}

	return -1
}
