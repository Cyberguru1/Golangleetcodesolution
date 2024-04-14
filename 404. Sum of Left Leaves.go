// Time C. O(n), Space C. O(1)
// runtime 2ms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) (res int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root != nil && root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			res += root.Left.Val
		}
		dfs(root.Left)
		dfs(root.Right)
		return
	}
	dfs(root)
	return
}

// more crazier(mature) soln <(-_-)>
// still 2ms though
func sumOfLeftLeaves(root *TreeNode) (res int) {
	var dfs func(root *TreeNode, check bool)
	dfs = func(root *TreeNode, check bool) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && check {
			res += root.Val
		}
		dfs(root.Left, true)
		dfs(root.Right, false)
		return
	}
	dfs(root, false)
	return
}