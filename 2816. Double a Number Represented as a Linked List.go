// Time C. O(n), Space C. O(1)
// runtime 2ms


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func doubleIt(head *ListNode) *ListNode {

	if head.Val > 4 {
		head = &ListNode{0, head}
	}
	trav := head
	for trav.Next != nil {
		trav.Val = (trav.Val * 2 ) % 10 
		if trav.Next.Val > 4 {
			trav.Val++
		}
		trav = trav.Next
	}

	trav.Val = (trav.Val * 2) % 10

	return head
}

