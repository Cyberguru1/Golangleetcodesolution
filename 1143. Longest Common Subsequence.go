// Time C. O(nm) Space C. O(nm), where n and m are length of text1 and text2 resp.
// runtime 219ms


func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1) + 1)
	for i:= len(text1); i > -1; i-- {
		dp[i] = make([]int, len(text2) + 1)
		for j := len(text2); j > -1; j--{
			if len(text1) == i || len(text2) == j{
				dp[i][j] = 0
			}else if (text1[i] == text2[j]){
				dp[i][j] = 1 + dp[i+1][j+1]
			}else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
