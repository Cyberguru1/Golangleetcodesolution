// runtime 0ms
// Time C. O(n), Space C. O(n)


func numDecodings(s string) int {
	dp := map[int]int{len(s):1}

	for i := len(s)-1; i > -1; i-- {
		if s[i] == '0'{
			dp[i] = 0
		} else {
			dp[i] = dp[i + 1]
		}
		fmt.Println(i)
		if ((i+1 < len(s)) && ((s[i] == '1') || ((s[i] == '2') && (stringInSlice(s[i+1]))))){
			dp[i] += dp[i + 2]
		}
	}

	return dp[0]
}

func stringInSlice(a byte) bool {
	list := []byte{'0', '1', '2', '3', '4', '5', '6'}
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}