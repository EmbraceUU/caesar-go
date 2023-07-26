package tree

// PostorderTraversal 【后序遍历 非递归】
// 左子树 -> 右子树 -> 根节点
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode

	for root != nil || len(stack) > 0 {
		// 先访问左子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		// 增加一个lastVisit, 确保右子树已经弹出, 再弹出根节点
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)

			lastVisit = node
		} else {
			root = node.Right
		}
	}

	return result
}

// PostorderTraversalII 后序遍历 迭代
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
func PostorderTraversalII(root *TreeNode) []int {
	// 准备两个栈
	// 一个栈用来遍历，一个栈用来存储结果
	// 先把root压栈
	// 开始出栈，出栈之后将左右两个节点入栈，然后将当前节点入结果栈
	// 周而复始

	if root == nil {
		return nil
	}

	var res []int
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		// 出栈当前节点
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 将左右两个子节点入栈
		// 先入左节点，再入右节点，出栈的顺序就是先右后左
		// 并且当前节点此时已经出栈，那么总体的出栈顺序就是 root -> right -> left
		// 最后再反转数组，得到一个遍历顺序为 left -> right -> root 的数组
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		// 将结果入栈
		res = append(res, cur.Val)
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
