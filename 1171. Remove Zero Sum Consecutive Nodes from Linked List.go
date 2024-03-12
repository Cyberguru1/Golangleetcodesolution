// Time C. O(n), Space C. O(n)
// runtime 5ms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode)( res *ListNode ){


	iter := head

	res = &ListNode{Val: 0, Next: head}
	mapp := map[int]*ListNode{0:res}

	pSum := 0
	for iter != nil {
		pSum += iter.Val
		mapp[pSum] = iter
		iter = iter.Next
	}

	pSum = 0
	iter = res

	for iter != nil {
		pSum += iter.Val
		iter.Next = mapp[pSum].Next
		iter = iter.Next
	}

	res = res.Next
	return
}