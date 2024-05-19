// Time C. O(), Space C. O()
//

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var totalSum int64
	count := 0
	positiveMin := math.MaxInt64
	negativeMax := math.MinInt64

	for _, nodeValue := range nums {
		nodeValAfterOperation := nodeValue ^ k
		totalSum += int64(nodeValue)
		netChange := nodeValAfterOperation - nodeValue

		if netChange > 0 {
			if netChange < positiveMin {
				positiveMin = netChange
			}
			totalSum += int64(netChange)
			count++
		} else {
			if netChange > negativeMax {
				negativeMax = netChange
			}
		}
	}

	if count%2 == 0 {
		return totalSum
	}
	return int64(math.Max(float64(totalSum-int64(positiveMin)), float64(totalSum+int64(negativeMax))))
}