package main

import (
	"math"
)

type Triplet struct {
	minSum       int
	secondMinSum int
	minSumIndex  int
}

func minFallingPathSum(grid [][]int) int {
	return minFallingPathSumHelper(0, grid).minSum
}
1289. Minimum Falling Path Sum II

func minFallingPathSumHelper(row int, grid [][]int) Triplet {
	if row == len(grid) {
		return Triplet{0, 0, 0}
	}

	nextRowTriplet := minFallingPathSumHelper(row+1, grid)
	currTrip := Triplet{math.MaxInt32, math.MaxInt32, -1}

	for col := 0; col < len(grid[0]); col++ {
		var value int
		if col != nextRowTriplet.minSumIndex {
			value = grid[row][col] + nextRowTriplet.minSum
		} else {
			value = grid[row][col] + nextRowTriplet.secondMinSum
		}

		if value <= currTrip.minSum {
			currTrip.secondMinSum = currTrip.minSum
			currTrip.minSum = value
			currTrip.minSumIndex = col
		} else if value < currTrip.secondMinSum {
			currTrip.secondMinSum = value
		}
	}

	return currTrip
}

