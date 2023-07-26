package tree

// InorderTraversal 中序遍历 迭代
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// 先把左边界压栈
	// 当左边界遍历完，开始出栈
	// 出栈的同时，访问节点
	// 并判断是否存在右节点，
	// 有则开始以右节点压左边界

	var res []int
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {

		// 先把当前root的左边界压栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 出栈操作
		if len(stack) > 0 {

			// 开始出栈，出栈后访问当前node
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res = append(res, node.Val)

			// 并判断是否有右节点
			// 如果有右节点，就把右节点给root
			// 然后在下一个循环中继续把新的左边界压栈
			if node.Right != nil {
				root = node.Right
			}
		}
	}
	return res
}
