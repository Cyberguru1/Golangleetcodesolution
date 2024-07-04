// Time C. O(n), Space C. O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func nodesBetweenCriticalPoints(head *ListNode) []int {
//     ress := make([]int, 0)

//     v := head.Next
//     prev := head
//     ind := 2

//     for v != nil {
//         if v.Next != nil {
//             if (v.Val > prev.Val && v.Val > v.Next.Val) ||  (v.Val < prev.Val && v.Val < v.Next.Val) {
//                 ress = append(ress, ind)
//             }
//         }
//         ind++
//         v = v.Next
//         prev = prev.Next
//     }

//     if len(ress) == 0 || len(ress) == 1 {
//         return []int{-1, -1}
//     }else if len(ress) == 2 {
//         val := ress[1] - ress[0]
//         return []int{val, val}
//     }else {
//         val1 := int(^uint(0)>>1)
//         prev := ress[0]
//         for _, k := range ress[1:] {
//             val1 = min(val1, k - prev)
//             prev = k
//         }
//         val2 := ress[len(ress)-1] - ress[0]
//         return []int{val1, val2}
//     }

//     return []int{}
// }

// second solution with no array | less memory

func nodesBetweenCriticalPoints(head *ListNode) (res []int) {
	res = []int{-1, -1}

	v := head.Next
	prev := head
	ind := 2
	firstCr, prevCr, minInd := 0, 0, int(^uint(0)>>1)

	for v != nil {
		if v.Next != nil {
			if (v.Val > prev.Val && v.Val > v.Next.Val) || (v.Val < prev.Val && v.Val < v.Next.Val) {
				if prevCr == 0 {
					firstCr, prevCr = ind, ind
				} else {
					minInd = min(minInd, ind-prevCr)
					prevCr = ind
				}
			}
		}
		ind++
		v = v.Next
		prev = prev.Next
	}

	if minInd != int(^uint(0)>>1) {
		res[0], res[1] = minInd, prevCr-firstCr
	}

	return
}