//Time C. O(n), Space C. O(n)
// runtime 3ms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {

	//initialize
	res := ""

	var dfs func(root *TreeNode, spath string)

	dfs = func(root *TreeNode, spath string) {

		if root == nil {
			return
		}

		spath += string(root.Val + 97)

		if root.Left == nil && root.Right == nil {
			spath = rev(spath)
			if res == "" {
				res = spath
			}
			cmp := strings.Compare(spath, res) // lexicographically compare
			if cmp < 0 {
				res = spath
			}
		}

		dfs(root.Left, spath)
		dfs(root.Right, spath)

		// backtrack
		spath = spath[:len(spath)-1]
	}
	dfs(root, "")

	return res
}

// can't believe there's no one liner to reverse a string in golang
//sigh <[^_-]>

func rev(str string) string {
	rres, n := []rune(str), len(str)
	for i := 0; i < n/2; i++ {
		rres[i], rres[n-1-i] = rres[n-1-i], rres[i]
	}
	return string(rres)
}
