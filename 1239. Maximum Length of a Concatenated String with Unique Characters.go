// Time C. O(n*n), Space C. O(n)
// runtime 9ms
func maxLength(arr []string) (res int) {
	check := func(s1 string) bool {
		mapp := map[rune]int{}
		for _, k := range s1 {
			if _, ok := mapp[k]; ok {
				return false
			}
			mapp[k] = 1
		}
		return true
	}
	var dfs func(path string, idx int)
	dfs = func(path string, idx int) {
		if check(path) {
			res = max(len(path), res)
		}
		if idx == len(arr) || !check(path) {
			return
		}
		for i := idx; i < len(arr); i++ {
			dfs(path+arr[i], i+1)
		}
	}
	dfs("", 0)
	return
}