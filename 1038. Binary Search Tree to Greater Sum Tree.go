// Time C. O(2^n), Space C. O(1)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {

	currSum := 0

	var trachg func(root *TreeNode)
	var sunchg func(root *TreeNode)

	trachg = func(root *TreeNode) {

		if root != nil {
			trachg(root.Left)
			currSum += root.Val
			trachg(root.Right)
		}
	}

	// accumulate currSum using inorder
	trachg(root)

	prev := 0

	sunchg = func(root *TreeNode) {

		if root != nil {
			sunchg(root.Left)
			currSum -= prev
			prev = root.Val
			root.Val = currSum
			sunchg(root.Right)
		}
	}

	// use same concept to subtract and change
	sunchg(root)

	return root
}
