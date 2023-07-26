package tree

// IsBalanced 平衡二叉树
// 分治法 + 递归
// 判断条件 左边平衡 && 右边平衡 && 两边高度差 <= 1
// 技巧: 返回数据二义性, 用-1表示不平衡, 用>=0表示高度
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	depth := isBalanced(root)
	if depth < 0 {
		return false
	}
	return true
}

func isBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dl := isBalanced(root.Left)
	dr := isBalanced(root.Right)

	if dl < 0 || dr < 0 {
		return -1
	}

	if dl-dr > 1 || dl-dr < -1 {
		return -1
	}

	if dl > dr {
		return dl + 1
	} else {
		return dr + 1
	}
}

// IsBalancedII 平衡二叉树
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
// https://leetcode-cn.com/problems/balanced-binary-tree/
func IsBalancedII(root *TreeNode) bool {
	isBFS, _ := isBalancedII(root)
	return isBFS
}

func isBalancedII(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftIsBalanced, leftHigh := isBalancedII(root.Left)
	rightIsBalanced, rightHigh := isBalancedII(root.Right)

	if !leftIsBalanced || !rightIsBalanced {
		return false, 0
	}

	if leftHigh-rightHigh > 1 || rightHigh-leftHigh > 1 {
		return false, 0
	}

	high := leftHigh
	if high < rightHigh {
		high = rightHigh
	}

	return true, high + 1
}
