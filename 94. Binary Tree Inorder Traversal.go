// iterative solution of inorder traversal
// Time C. O(n) Space C. O(n)


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {

	stackLike := []*TreeNode{}
	ind := -1

	curr := root
	result := []int{}

	for ((curr != nil) || (ind != -1)){

		for (curr != nil){
			stackLike = append(stackLike, curr)
			ind += 1
			curr = curr.Left
		}

		curr = stackLike[ind]
		result = append(result, curr.Val)
		stackLike = append(stackLike[:ind], stackLike[ind+1:]...)
		ind--

		curr = curr.Right


	}
	return result
}