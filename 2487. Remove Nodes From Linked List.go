// Time C. O(n), Space C. O(N)
// Runtime 6ms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	stack := []int{}

	trav := head

	for trav != nil {
		stack = append(stack, trav.Val)
		trav = trav.Next
	}

	nStack := []int{}
	prev := 0

	for i := len(stack) - 1; i >= 0; i-- {
		if prev <= stack[i] {
			prev = stack[i]
			nStack = append(nStack, prev)
		}
	}

	nHead := &ListNode{0, nil}
	curr := nHead

	for i := len(nStack) - 1; i >= 0; i-- {
		temp := &ListNode{nStack[i], nil}
		curr.Next = temp
		curr = temp

	}

	return nHead.Next

}