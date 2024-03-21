// Time C. O(n + m), Space C. O(1) where n and m are size of list1 and list2 resp
// runtime

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	l = list1
	for i := 0; i < a - 1; i++{
		l = l.Next
	}

	r, prev := l, l

	for i := 0; i < b + 1; i++ {
		r = r.Next
		prev.Next = nil
		prev = r
	}

	l2tail := list2

	for l2tail.Next != nil {
		l2tail = l2tail.Next
	}

	l.Next = list2
	l2tail.Next = r

	return list1
}