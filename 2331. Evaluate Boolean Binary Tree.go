/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution 1 -- using stack

func evaluateTree(root *TreeNode) bool {

    sigMap := map[int]bool{0:false, 1:true}
    revSigMap := map[bool]int{false:0, true:1}

    stack := make([]int, 0)
    var trav func(root *TreeNode)

    trav = func(root *TreeNode){
        if root != nil {
            trav(root.Left)
            trav(root.Right)
            if (root.Val == 3 || root.Val == 2) {
                f := stack[len(stack)-1]
                l := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                if root.Val == 3 {
                    stack = append(stack, revSigMap[(sigMap[f] && sigMap[l])])
                }else{
                    stack = append(stack, revSigMap[(sigMap[f] || sigMap[l])])
                }
            }else{
                stack = append(stack, root.Val)
            }
        }
    }

    trav(root)
    return sigMap[stack[len(stack)-1]]

}

// fu**k
// solution 2 -- pure recursion

func evaluateTree(root *TreeNode) bool {

	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}

	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}

	return evaluateTree(root.Right) && evaluateTree(root.Left)
}
