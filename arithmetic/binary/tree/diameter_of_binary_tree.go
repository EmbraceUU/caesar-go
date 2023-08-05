package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	// 后序遍历
	// 左子树提供一个最远距离
	// 右子树提供一个最远距离
	// 取左右子树中较长的一个 + 当前路径，返回
	result := new(int)
	var maxDiameter func(root *TreeNode, result *int) int
	maxDiameter = func(root *TreeNode, result *int) int {
		if root == nil {
			return 0
		}

		leftMax := maxDiameter(root.Left, result)
		rightMax := maxDiameter(root.Right, result)

		if *result < leftMax+rightMax {
			*result = leftMax + rightMax
		}
		if leftMax > rightMax {
			return leftMax + 1
		} else {
			return rightMax + 1
		}
	}

	maxDiameter(root, result)

	return *result
}
