/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {

	Minl, Maxl := root.Val, root.Val
	const INT_MAXl = int(^uint(0) >> 1)

	ans := -(INT_MAXl - 1)

	var dfs func(root *TreeNode, Maxl, Minl int, ans *int)

	dfs = func(root *TreeNode, Maxl, Minl int, ans *int) {
		if root == nil {
			return
		}

		*ans = max(*ans, Maxl - root.Val, root.Val - Minl)


		if Maxl <= root.Val { Maxl = root.Val }
		if Minl >= root.Val { Minl = root.Val }

		dfs(root.Left, Maxl, Minl, ans)
		dfs(root.Right, Maxl, Minl, ans)
	}

	dfs(root, Maxl, Minl, &ans)

	return ans

}

