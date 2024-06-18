// first approach -- more longer

// type mapping struct {
//     diff int
//     prof int
// }

// func maxProfitAssignment(difficulty []int, profit []int, worker []int) (res int ){

//     pdMap := []mapping{}
//     dpMap := []mapping{}

//     for i, val := range difficulty {
//         pdMap = append(pdMap, mapping{ val, profit[i] })
//         dpMap = append(dpMap, mapping{ val, profit[i]})
//     }

//     sort.Slice(pdMap, func(i, j int)bool {
//         return pdMap[i].prof > pdMap[j].prof
//     })
//     sort.Slice(dpMap, func(i, j int)bool {
//         return dpMap[i].diff < dpMap[j].diff
//     })

//     t, d := 0, len(dpMap) - 1
//     for t < len(dpMap) && d >= 0 {
//         if dpMap[d].diff >= pdMap[t].diff && dpMap[d].prof <= pdMap[t].prof {
//             dpMap[d] = mapping{ dpMap[d].diff, pdMap[t].prof}
//             d--
//             continue
//         }
//         t++
//     }

//     for len(worker) > 0 {
//         curr := worker[len(worker)-1]
//         worker = worker[:len(worker)-1]

//         for i := len(dpMap) - 1; i >= 0; i-- {
//             if curr >= dpMap[i].diff {
//                 res += dpMap[i].prof
//                 break
//             }
//         }
//     }

//     return
// }

// second approach -- involving sorting worker -- more easier and faster


type mapping struct {
	diff int
	prof int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) (res int ){

	pdMap := []mapping{}

	for i, val := range difficulty {
		pdMap = append(pdMap, mapping{ val, profit[i] })
	}

	sort.Slice(pdMap, func(i, j int)bool {
		return pdMap[i].diff < pdMap[j].diff
	})

	sort.Ints(worker)

	maxwork, j := 0, 0

	for i := 0; i < len(worker); i++ {
		// get max profit in loop range
		for j < len(difficulty) && worker[i] >= pdMap[j].diff {
			maxwork = max(maxwork, pdMap[j].prof)
			j++
		}

		res += maxwork
	}

	return
}