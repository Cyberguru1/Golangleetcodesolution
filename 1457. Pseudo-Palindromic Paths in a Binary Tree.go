// Time C. O(n), Space C. O(n)---> due to runtime stack
// Runtime 212ms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pseudoPalindromicPaths(root *TreeNode) int {

	res, path := 0, 0

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root != nil {
			path = path ^ (1 << root.Val)
			if root.Left == nil && root.Right == nil {
				if (path & (path - 1)) == 0 {
					res++
				}
			}
			dfs(root.Left)
			dfs(root.Right)
			path = path ^ (1 << root.Val)
		}
	}

	dfs(root)

	return res
}
