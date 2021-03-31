package binary_tree

import (
	"fmt"
)

// 以根访问顺序决定是什么遍历
// 左子树都是优先右子树

// 前序遍历 - 递归
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 访问当前根节点
	fmt.Println(root.Val)
	// 前序遍历左子树
	preorderTraversal(root.Left)
	// 前序遍历右子树
	preorderTraversal(root.Right)
}

// 前序遍历 - 非递归
func preorderTraversal2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
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

// 中序遍历 - 非递归
// 左子树 -> 根节点 -> 右子树
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		// 一直向左压栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 当到达左子树尽头, 开始pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 打印节点数据
		fmt.Println(node.Val)
		// 转移到右子树, 再中序遍历
		root = node.Right
	}
}

// 后序遍历
// 左子树 -> 右子树 -> 根节点
func postorderTraversal(root *TreeNode) []int {
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

// DFS 深度遍历 从下到上 分治法
func preorderTraversal3(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
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
