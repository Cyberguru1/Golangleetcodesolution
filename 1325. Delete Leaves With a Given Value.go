// TIme C. O(2^n), Space C. O(n)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func removeLeafNodes(root *TreeNode, target int) *TreeNode {

	var traverse func(root *TreeNode) *TreeNode

	traverse = func(root *TreeNode) *TreeNode {

		if root == nil {
			return nil
		}

		root.Left = traverse(root.Left)
		root.Right = traverse(root.Right)


		if root.Left == nil && root.Right == nil && root.Val == target {
			return nil
		}

		return root
	}

	return traverse(root)
}
