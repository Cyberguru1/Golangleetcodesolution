// MR Beautiful
// 

func beautifulSubsets(nums []int, k int) (count int) {

	isbeauty := func(entry []int) bool {
		if len(entry) == 1 {
			return true
		}

		for i := 0; i < len(entry); i++ {
			for j := 0; j < len(entry); j++ {
				if j == i {
					continue
				}
				if (entry[i] - entry[j]) == k {
					return false
				}
			}
		}

		return true
	}

	var subsets func(arr, subs []int, ind int)

	subsets = func(arr, subs []int, ind int) {
		if len(subs) != 0 {
			if isbeauty(subs) {
				count++
			}
		}
		for j := ind; j < len(arr); j++ {
			subs = append(subs, arr[j])
			subsets(arr, subs, j+1)
			subs = subs[:len(subs)-1]
		}
	}

	subsets(nums, []int{}, 0)

	return

}