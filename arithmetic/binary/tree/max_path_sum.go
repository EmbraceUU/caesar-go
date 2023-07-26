package tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxSum = math.MinInt64

func maxPathSum(root *TreeNode) int {
	maxGain(root)
	return maxSum
}

func maxGain(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 左子树贡献值
	left := max(maxGain(root.Left), 0)
	// 右子树贡献值
	right := max(maxGain(root.Right), 0)

	// 更新maxSum
	current := left + right + root.Val
	maxSum = max(maxSum, current)

	// 返回当前节点贡献值
	return root.Val + max(left, right)
}
