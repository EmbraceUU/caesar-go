package tree

// maxDepth 借助【递归】和【深度优先遍历】
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)

	return max(ld+1, rd+1)
}
