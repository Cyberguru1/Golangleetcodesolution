// First Implementation Using Arrays
// Time C. O(n), Space C. O(n)
// runtime 2ms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	con := []*ListNode{}

	iter := head

	for iter != nil {
		con = append(con, iter)
		iter = iter.Next
	}

	link := con[len(con)-1]

	for i := len(con) - 2; i > -1; i-- {
		link.Next = con[i]
		link = con[i]
	}

	con[0].Next = nil

	return con[len(con)-1]
}

// Second Implementation Without Arrays, and constant space
// Time C. O(n), Space C. O(1)
// runtime 0ms

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode = nil
	next := head

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

