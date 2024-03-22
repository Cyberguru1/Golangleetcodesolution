// Time C. O(n), Space C. O(n)
// runtime 117ms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

	con := []*ListNode{}
	iter := head

	for iter != nil {
		con = append(con, iter)
		iter = iter.Next
	}

	l := 0
	r := len(con) - 1

	for l <= r {
		if con[l].Val != con[r].Val {
			return false
		}

		l++
		r--
	}

	return true
}