// Iime C. O(n^2) Space C. O(n)
// runtime 0ms

func removeKdigits(num string, k int) (string) {

	res := []rune{}

	for _, c := range num {
		for len(res) > 0 && res[len(res) - 1] > c && k > 0 {
			res = res[:len(res) - 1]
			k--
		}

		if len(res) > 0 || c != '0'{
			res = append(res, c)
		}
	}


	for len(res) > 0 && k > 0 {
		res = res[:len(res) - 1]
		k--
	}

	if len(res) == 0 {
		return "0"
	}

	return string(res)

}