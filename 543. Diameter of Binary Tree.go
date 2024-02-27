/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {

	 var trav func(root *TreeNode) int

	 res := 0
	 trav = func(root *TreeNode) int {
		 if root != nil {
			 left := trav(root.Left)
			 right := trav(root.Right)
			 if left+right > res {
				 res = left+right
			 }
			 return max(left, right) + 1
		 }
		 return 0
	 }

	 trav(root)
	 return res
}
