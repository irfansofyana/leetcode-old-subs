/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(l *TreeNode, r *TreeNode) bool {
	if l == nil && r != nil {
		return false
	}
	if l != nil && r == nil {
		return false
	}
	if l == nil && r == nil {
		return true
	}
	if l.Val != r.Val {
		return false
	}
	return helper(l.Left, r.Right) && helper(l.Right, r.Left)
}