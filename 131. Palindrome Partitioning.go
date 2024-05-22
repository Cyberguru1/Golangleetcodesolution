// Time C. O(2^n), Space C. O(n)

func partition(s string) (result [][]string) {

	ispal := func(str string) bool {
		newstr := ""

		for _, k := range str {
			newstr = string(k) + newstr
		}

		if newstr != str {
			return false
		}

		return true
	}

	var backtrack func(st int, path []string)

	backtrack = func(st int, path []string) {
		if st == len(s){
			result = append(result, append([]string{}, path...))
			return
		}

		for j := st+1; j <= len(s); j++ {
			if ispal(s[st:j]){
				backtrack(j, append(path, s[st:j]))
			}
		}
	}

	backtrack(0, []string{})

	return
}
