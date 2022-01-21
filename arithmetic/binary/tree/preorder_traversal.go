package tree

import "fmt"

// PreorderTraversal 【前序遍历 递归】
func PreorderTraversal(root *Node) {
	if root == nil {
		return
	}
	// 访问当前根节点
	fmt.Println(root.Val)
	// 前序遍历左子树
	PreorderTraversal(root.Left)
	// 前序遍历右子树
	PreorderTraversal(root.Right)
}

// PreorderTraversalII 【前序遍历 非递归】
func PreorderTraversalII(root *Node) {
	if root == nil {
		return
	}
	stack := make([]*Node, 0)
	for root != nil || len(stack) >= 0 {
		for root != nil {
			// 访问当前节点
			fmt.Println(root.Val)
			stack = append(stack, root)
			// 转移到左子树
			root = root.Left
			// 到叶子节点后退出
		}
		// LIFO pop出根节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 开始前序遍历右子树
		root = node.Right
	}
}

// PreorderTraversalIII 【DFS 深度遍历 分治法】
// DFS中 关键点是【递归以及回溯】
// 采用递归的方式, 先分别递归返回结果, 然后合并结果, 合并的时候, 可以控制访问的顺序
// 可以简单的实现 preorder/inorder/postorder 这三种顺序
func PreorderTraversalIII(root *Node) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *Node) []int {
	var result []int
	if root == nil {
		return nil
	}

	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// PreorderTraversalIV 前序遍历，迭代
// 左神的思路
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func PreorderTraversalIV(root *Node) []int {
	var res []int

	if root == nil {
		return res
	}

	// 如果想非递归的方式遍历二叉树，
	// 就需要借助栈来代替递归中的调用栈，
	// 先把root压栈
	stack := make([]*Node, 0)
	stack = append(stack, root)

	for len(stack) > 0 {

		// 先序遍历，先访问根节点
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 弹出栈时打印
		res = append(res, cur.Val)

		// 先把右节点压栈
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		// 再把左节点压栈
		// 这个顺序能保证先出栈的是左节点，然后是右节点
		// 正好符合前序遍历中的 root -> left -> right 顺序
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return res
}
