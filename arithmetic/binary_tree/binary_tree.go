package binary_tree

import (
	"fmt"
)

// 以根访问顺序决定是什么遍历
// 左子树都是优先右子树

// PreorderTraversal 【前序遍历 递归】
func PreorderTraversal(root *TreeNode) {
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
func PreorderTraversalII(root *TreeNode) {
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

// InorderTraversal 【中序遍历 非递归】
// 左子树 -> 根节点 -> 右子树
func InorderTraversal(root *TreeNode) {
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

// --------------------------- 分界线 --------------------------- //

// PreorderTraversalIII 【DFS 深度遍历 分治法】
// DFS中 关键点是【递归以及回溯】
// 采用递归的方式, 先分别递归返回结果, 然后合并结果, 合并的时候, 可以控制访问的顺序
// 可以简单的实现 preorder/inorder/postorder 这三种顺序
func PreorderTraversalIII(root *TreeNode) []int {
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

// --------------------------- 分界线 --------------------------- //

// LevelOrder 【BFS 广度优先搜索 第一版】
func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		var level []int
		var queueTemp []*TreeNode
		for i := range queue {
			level = append(level, queue[i].Val)
			if queue[i].Left != nil {
				queueTemp = append(queueTemp, queue[i].Left)
			}
			if queue[i].Right != nil {
				queueTemp = append(queueTemp, queue[i].Right)
			}
		}
		result = append(result, level)
		queue = queueTemp
	}

	return result
}

// LevelOrderII 【BFS 广度优先搜索 改进版】
// 遍历当前层，并标记下一层, 然后再遍历下一层, 以此循环遍历
// TODO 为什么改进版比第一版的内存消耗更大了 ？
func LevelOrderII(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// 遍历当前层
	for len(queue) > 0 {
		// 当前level遍历的结果
		list := make([]int, 0)
		// 先得到当前queue的长度，遍历当前层，再添加下一层
		l := len(queue)
		// 只遍历预先固定的长度, 这样不需要再每次创建queueTemp~~
		for i := 0; i < l; i++ {
			// 每次将队首出列, 先进先出, 这样即使queue变化了, 也不影响下面的继续访问
			level := queue[0]
			queue = queue[1:]

			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		// 遍历完当前层, 将结果集合并到result中
		result = append(result, list)
	}
	return result
}

// --------------------------- 分界线 --------------------------- //

// MaxDepth 求二叉树的最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dl := MaxDepth(root.Left)
	dr := MaxDepth(root.Right)

	if dl > dr {
		return dl + 1
	} else {
		return dr + 1
	}
}

// --------------------------- 分界线 --------------------------- //

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
