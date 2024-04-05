// Time C. O(n), Space C. O(1)
// Best Runtime 0ms

func makeGood(s string) (res string) {
	res = s
	for true {
		change := false
		for i := 0; i < len(res)-1; i++ {
			if rune(res[i]) >= 65 && rune(res[i]) < 97 || rune(res[i+1]) >= 65 && rune(res[i+1]) < 97 {
				if rune(res[i])+32 == rune(res[i+1]) {
					res = res[:i] + res[i+2:]
					change = true
				} else if rune(res[i])-32 == rune(res[i+1]) {
					res = res[:i] + res[i+2:]
					change = true
				}
			}
		}
		if !change {
			break
		}
	}
	return
}