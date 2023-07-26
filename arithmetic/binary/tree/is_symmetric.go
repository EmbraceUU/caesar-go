package tree

// isSymmetric 对称二叉树，不能用【迭代】的方式进行中序遍历得出，这样只能得到一个非空的序列
// https://leetcode-cn.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	// 使用【递归】的方式
	// 自己的左子树 == 自己的右子树
	var symmetric func(n1 *TreeNode, n2 *TreeNode) bool
	symmetric = func(n1 *TreeNode, n2 *TreeNode) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}

		return symmetric(n1.Left, n2.Right) && symmetric(n1.Right, n2.Left)
	}
	return symmetric(root.Left, root.Right)
}

// isSymmetricIterative 【迭代】，同样需要对比自己的左和右
func isSymmetricIterative(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left, root.Right)

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i = i + 2 {
			n1 := queue[0]
			queue = queue[1:]
			n2 := queue[0]
			queue = queue[1:]

			if n1 == nil && n2 == nil {
				continue
			}
			if n1 == nil || n2 == nil {
				return false
			}
			if n1.Val != n2.Val {
				return false
			}

			queue = append(queue, n1.Left, n2.Right, n1.Right, n2.Left)
		}
	}
	return true
}
