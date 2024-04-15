// Time C. O(N) where n is the number of root to nodes, Space C. O(N)
// Runtime 3ms
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) (res int) {

	paths := ""

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {

		if root == nil {
			return
		}

		paths = fmt.Sprintf("%s%s", paths, strconv.Itoa(root.Val))
		fmt.Println(paths)

		if root.Left == nil && root.Right == nil {
			s, _ := strconv.Atoi(paths)
			res += s
			paths = paths[:len(paths)-1]
			return
		}

		dfs(root.Left)
		dfs(root.Right)
		paths = paths[:len(paths)-1]
		return
	}
	dfs(root)
	return
}