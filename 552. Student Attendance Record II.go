// Clearly not my solution
// still learning the ropes of using arrays for dp problems

func checkRecord(n int) int {
	const MOD = 1000000007

	// Handle small values of n directly
	switch n {
	case 1:
		return 3
	case 2:
		return 8
	case 3:
		return 19
	}

	// dp[i] represents the count of valid sequences of length i without 'A'
	dp := make([]int, n+1)
	dp[0] = 1 // Base case: an empty sequence
	dp[1] = 2 // "P", "L"
	dp[2] = 4 // "PP", "PL", "LP", "LL"
	dp[3] = 7 // "PPP", "PPL", "PLP", "PLL", "LPP", "LPL", "LLP"

	// Fill the dp array for lengths from 4 to n
	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % MOD
	}

	// Calculate the number of valid sequences with exactly one 'A'
	withA := 0
	for i := 0; i < n; i++ {
		withA = (withA + dp[i]*dp[n-1-i]) % MOD
	}

	// The total count is the sum of sequences with and without 'A'
	return (dp[n] + withA) % MOD
}
