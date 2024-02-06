// Time C. O(n*m) Space C. O(n) where n is len of strs array and m len of each element of the array
// runtime 21ms

func groupAnagrams(strs []string) (res [][]string) {
	mapp := map[string][]string{}


	sort_func := func(str1 string) string {

		unsorted := []rune(str1)

		sort.Slice(unsorted, func(i, j int) bool {
			return unsorted[i] < unsorted[j]
		})

		return string(unsorted)
	}

	for _, k := range strs {
		sorted := sort_func(k)

		mapp[sorted] = append(mapp[sorted], k)
	}

	for _, k := range mapp {
		res = append(res, k)
	}

	return

}