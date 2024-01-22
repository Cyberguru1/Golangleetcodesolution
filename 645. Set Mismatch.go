// Time C. O(n), Space C. O(n)
// runtime 38ms

func findErrorNums(nums []int) (res []int) {

	mapp := map[int]int{}

	for _, k := range nums {
		if _, ok := mapp[k]; ok {
			res = append(res, k)
		}
		mapp[k]++
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := mapp[i+1]; !ok {
			res = append(res, i+1)
		}
	}
	return
}