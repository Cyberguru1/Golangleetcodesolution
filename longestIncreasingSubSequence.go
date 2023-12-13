package main

import (
	"log"
)


func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = 1
	}

	for i :=1; i < len(nums); i++ {
		for j :=0; j < i; j++ {
			if nums[i] > nums[j]{
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	} 
	return slices.Max(dp)
}

func main(){

	arr := []int{10,9,2,5,3,7,101,18}
	result := lengthOfLIS(arr)

	log.Println(result)
	
}