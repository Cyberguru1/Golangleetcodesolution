import "slices"

func longestCommonPrefix(strs []string) string {

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(strs, lenCmp)

	result := ""

	for i := 0; i < len(strs[0]); i++ {
		good := true
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				good = false
			}
		}

		if good {
			result = fmt.Sprintf("%s%s", string(result), string(strs[0][i]))
		} else {
			break
		}
	}

	return result
}