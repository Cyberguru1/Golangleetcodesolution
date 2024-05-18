// Time C. O(2^n), Space C. O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {

	count := 0

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		t := root.Val + l + r - 1
		if t < 0 {
			count += (t * -1)
		}else {
			count += t
		}
		return t
	}

	dfs(root)
	return count
}