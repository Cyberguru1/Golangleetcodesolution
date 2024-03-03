// Enjoyed this, been long i sovled a linked list problem
// Runtime 2ms
// Time C. O(n), Space C. O(n)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	arr := []*ListNode{}

	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	pos := len(arr) - n

	head = arr[0]
	prev := head

	if pos == 0 {
		return head.Next
	}

	for i:= 1; i < len(arr); i++ {
		if i == pos {
			prev.Next = nil
			continue
		}
		prev.Next = arr[i]
		prev = arr[i]
	}

	return head

}