// Time C. O(n), Space C. O(n)
// Runtime 10ms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {

	out := []*TreeNode{root}

	leftmost := 0

	for len(out) > 0 {
		last := out[0]
		out = out[1:]

		leftmost = last.Val

		if last.Right != nil {
			out = append(out, last.Right)
		}
		if last.Left != nil {
			out = append(out, last.Left)
		}
	}
	return leftmost
}