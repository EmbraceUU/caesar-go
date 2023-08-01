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

var res int
var depth int

func maxDepthII(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxD(root)
	return res
}

func maxD(root *TreeNode) {
	if root == nil {
		return
	}

	depth++
	if root.Left == nil && root.Right == nil {
		res = max(res, depth)
	} else {
		maxD(root.Left)
		maxD(root.Right)
	}
	depth--
}
