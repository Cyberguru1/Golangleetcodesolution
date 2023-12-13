package main

import (
	"math"
	"fmt"
)

func coinChange(coins []int, amount int) float64 {
	dp := make([]float64, amount + 1)
	for i := range dp {
		dp[i] = math.Inf(1)
	}

	dp[0] = 0

	for _ , coin := range coins {
		for i := coin; i < amount + 1; i++ {
			dp[i] = math.Min(dp[i], dp[i-coin]+1)
		}
	}
	if math.IsInf(dp[amount], 1) {
		return -1
	}else{
		return dp[amount]
	}
}




// func coinChange(coins []int, amount int) int {
// 	dp := make([]float64, amount + 1)
// 	for i := range dp {
// 		dp[i] = math.Inf(1)
// 	}

// 	dp[0] = 0

// 	for _ , coin := range coins {
// 		for i := coin; i < amount + 1; i++ {
// 			dp[i] = math.Min(dp[i], dp[i-coin]+1)
// 		}
// 	}
// 	if math.IsInf(dp[amount], 1) {
// 		return -1
// 	}else{
// 		return int(dp[amount])
// 	}
// }

func main(){
	result := coinChange([]int{1,2,5}, 11)
	fmt.Println("Result: ",result)
}
