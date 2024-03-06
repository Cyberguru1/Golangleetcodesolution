// Time C. O(n), Space C. O(1)
// best runtime 4ms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	tort := head
	hare := head.Next

	for tort != nil && hare != nil {
		if hare == tort {
			return true
		}
		tort = tort.Next
		if hare.Next == nil {
			return false
		}else {
			hare = hare.Next.Next // two times faster than tortise
		}
	}
	return false
}