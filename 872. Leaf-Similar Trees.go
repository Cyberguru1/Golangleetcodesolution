// first solution
// Time C. O(n), Space C. O(m) where m is the number of leafnodes

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	queue1 := []int{}
	queue2 := []int{}
	var extractLeaf func(root *TreeNode, queue *[]int)
	extractLeaf = func(root *TreeNode, queue *[]int) {
		if root == nil {
			return
		}
		if (root.Left == nil) && (root.Right == nil) {
			*queue = append(*queue, root.Val)
		}
		extractLeaf(root.Left, queue)
		extractLeaf(root.Right, queue)
	}
	extractLeaf(root1, &queue1)
	extractLeaf(root2, &queue2)

	if len(queue1) != len(queue2) {
		return false
	}
	for i, _ := range queue1 {

		if queue1[i] != queue2[i] {
			return false
		}
	}
	return true
}