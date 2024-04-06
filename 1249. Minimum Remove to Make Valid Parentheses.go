// Time C. O(n), Space C. O(n)
// Runtime 1337...to the best i could optimize

func minRemoveToMakeValid(s string) string {
	sarray := []int{}
	mapp := map[int]rune{}

	for i, v := range s {
		if v == '(' {
			sarray = append(sarray, i)
		} else if v == ')' && len(sarray) > 0 {
			sarray = sarray[:len(sarray)-1]
		} else if v == ')' {
			continue
		}
		mapp[i] = v
	}
	res := ""

	for i := 0; i < len(s); i++ {
		if len(sarray) > 0 && i == sarray[0] {
			sarray = sarray[1:]
			continue
		}
		if k, ok := mapp[i]; ok {
			res += string(k)
		}
	}
	return res
}
