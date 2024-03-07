// Time C. O(n), Space C. O(n)
// Best runtime 0ms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	nodes := []*ListNode{}

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	size := len(nodes)

	return nodes[size/2]
}
