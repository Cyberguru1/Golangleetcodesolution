// Time C. O(n), Space C. O(n)
// runtime 2ms

type void struct {}
func uniqueOccurrences(arr []int) bool {

	mapp := map[int]int{}

	for _, k := range arr {
		mapp[k]++
	}
	var mem void
	check := map[int]void{}

	for _, k := range mapp {
		if _, ok := check[k]; ok {
			return false
		}
		check[k] = mem
	}

	return true

}