// Time C. O(n), Space C. O(n)
// runtime 10ms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {

	con := []*ListNode{}

	iter := head

	for iter != nil {
		con = append(con, iter)
		iter = iter.Next
	}

	tlen := len(con)
	l := tlen - 1

	for i := 0; i < tlen / 2; i++ {
		con[i].Next = con[l]
		if i != tlen - 1 {
			con[l].Next = con[i+1]
		}
		l--
	}

	con[tlen/2].Next = nil
}