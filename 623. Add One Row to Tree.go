// Time C. O(n) where n is the desired depth, Space C. O(1)
// Runtime 3ms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode { 
	var dfs func(root *TreeNode, depth int) 
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == 1 {
			p1 := root.Left
			p2 := root.Right
			root.Left = &TreeNode{val, p1, nil}
			root.Right = &TreeNode{val, nil, p2}
			return // return and don't go in further !!
		}
		dfs(root.Left, depth-1)
		dfs(root.Right, depth-1)
	}

	if depth - 1 == 0 {
		return &TreeNode{val, root, nil}
	}
	dfs(root, depth - 1)
	return root
}
