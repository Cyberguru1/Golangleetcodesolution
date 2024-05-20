// Time c> O(n), Space C. O(n)

func subsetXORSum(nums []int) int {
	subb := make([][]int, 0)

	var subset func(arr []int, i int)

	subset = func(arr []int, i int){
		subb = append(subb, append([]int{}, arr...))

		for j := i; j < len(nums); j++ {
			arr = append(arr, nums[j])
			subset(arr, j+1)
			arr = arr[:len(arr)-1]
		}
	}

	subset([]int{}, 0)

	tsum := 0

	for _, k := range subb[1:] {
		if len(k) < 1 {
			tsum += k[0]
		}
		pxum := 0

		for _, v := range k {
			pxum ^= v
		}
		tsum+= pxum
	}
	return tsum
}
