
// func findCenter(edges [][]int) (res int){

//     countMap := map[int]int{}
//     maxx := 0
//     for _, con := range edges {
//         for _, val := range con {
//             countMap[val]++
//             if countMap[val] > maxx {
//                 maxx = countMap[val]
//                 res = val
//             }
//         }
//     }
//     return
// }

// or more better sol

func findCenter(edges [][]int) int{

	countMap := map[int]int{}

	for _, con := range edges {
		countMap[con[0]]++
		countMap[con[1]]++

		if countMap[con[0]] > 1 {
			return con[0]
		}

		if countMap[con[1]] > 1 {
			return con[1]
		}
	}
	return -1
}