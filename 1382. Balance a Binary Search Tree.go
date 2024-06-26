// Time C. O(logn), Space C. O(1)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func balanceBST(root *TreeNode) *TreeNode {

	arr := []*TreeNode{}

	var traverse func (root *TreeNode)
	var buildBst func (start, end int) *TreeNode

	traverse = func(root *TreeNode) {
		if root != nil {
			traverse(root.Left)
			arr = append(arr, root)
			traverse(root.Right)
		}
	}

	traverse(root)

	// use binary search approach to populate a new tree
	// keep spliting...

	buildBst = func(start, end int) *TreeNode {

		if start > end {
			return nil
		}

		mid := (start + end) / 2
		d := arr[mid]

		d.Left = buildBst(start, mid-1)
		d.Right = buildBst(mid+1, end)

		return d
	}

	return buildBst(0, len(arr)-1)
}