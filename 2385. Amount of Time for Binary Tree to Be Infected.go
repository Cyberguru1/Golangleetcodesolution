/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func amountOfTime(root *TreeNode, start int) int {

	var maxDistance int

	var traverse func(root *TreeNode, start int) int

	traverse = func(root *TreeNode, start int) int {
		depth := 0

		if root == nil {
			return depth
		}

		leftDepth := traverse(root.Left, start)
		rightDepth := traverse(root.Right, start)

		if root.Val == start {
			maxDistance = max(leftDepth, rightDepth)
			depth = -1
		} else if leftDepth >= 0 && rightDepth >= 0 {
			depth = max(leftDepth, rightDepth) + 1
		} else {
			distance := Abs(leftDepth) + Abs(rightDepth)
			maxDistance = max(maxDistance, distance)
			depth = min(leftDepth, rightDepth) - 1
		}

		return depth
	}

	traverse(root, start)

	return maxDistance

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}