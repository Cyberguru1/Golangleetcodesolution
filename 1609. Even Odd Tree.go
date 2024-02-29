// Enjoyed the mental stress while solving this (medium) challenge
// Time C. O(n) Space C. O(n)
// Best runtime 167ms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	mapper := map[int][]int{}
	var trav func(root *TreeNode, lvl int) bool
	trav = func(root *TreeNode, lvl int) bool {
		if root != nil {
			val := -1
			if len(mapper[lvl]) > 0 {
				val = mapper[lvl][len(mapper[lvl])-1]
			}
			if lvl%2 == 0 {
				if root.Val%2 == 0 || (val >= root.Val && val != -1) {
					return false
				}
				mapper[lvl] = append(mapper[lvl], root.Val)
			} else {
				if root.Val%2 != 0 || (val <= root.Val && val != -1) {
					return false
				}
				mapper[lvl] = append(mapper[lvl], root.Val)
			}
			return trav(root.Left, lvl+1) && trav(root.Right, lvl+1)
		}
		return true
	}
	return trav(root, 0)
}