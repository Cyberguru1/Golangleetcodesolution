/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

	var traverse func(p *TreeNode, q *TreeNode) bool

	res := true

	traverse = func(p, q *TreeNode) bool {
		if p != nil && q != nil {
			if p.Val != q.Val{
				res = res && false
			}
			traverse(p.Left, q.Left)
			traverse(p.Right, q.Right)
		}else if (p == nil && q != nil )|| ( q == nil && p != nil ){
			res = res && false
		}
		return res
	}

	return traverse(p, q)

}100. Same Tree
