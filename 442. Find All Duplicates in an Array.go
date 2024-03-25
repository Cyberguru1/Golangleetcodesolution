// Time C. O(n), Space C. O(1)
// runtime 50ms

type void struct {
}

func findDuplicates(nums []int) (out []int) {

	var con void

	mapp := map[int]void{}

	for _, i := range nums {

		if _, ok := mapp[i]; ok {
			out = append(out, i)
			continue
		}
		mapp[i] = con
	}

	return

} /*  */