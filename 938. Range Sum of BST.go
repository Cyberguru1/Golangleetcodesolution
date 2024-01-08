// Time C. O(n), Space C. O(1)
// Runtime 70ms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {

	sum := 0

	var traverseInorder func(root *TreeNode)

	traverseInorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverseInorder(root.Left)
		if (root.Val >= low) && (root.Val <= high) {
			sum += root.Val
		}
		if root.Val > high {
			return
		}
		traverseInorder(root.Right)
	}

	traverseInorder(root)

	return sum

}