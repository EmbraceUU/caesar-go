package tree

// LevelOrder 二叉树层次遍历
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
// （即逐层地，从左到右访问所有节点）。
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func LevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	// 层次遍历要借助队列 queue
	// 把根节点入队列
	// 然后将队列清空的同时，将出队列的节点的左右节点入队列

	var res [][]int
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level := make([]int, 0)
		queueLevel := make([]*Node, 0)

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queueLevel = append(queueLevel, cur.Left)
			}
			if cur.Right != nil {
				queueLevel = append(queueLevel, cur.Right)
			}
		}

		res = append(res, level)
		queue = queueLevel

	}

	return res
}

// LevelOrderII BFS 改进版，但是内存消耗变大了
func LevelOrderII(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	// 遍历当前层
	for len(queue) > 0 {
		// 当前level遍历的结果
		list := make([]int, 0)
		// 先得到当前queue的长度，遍历当前层，再添加下一层
		l := len(queue)
		// 只遍历预先固定的长度, 这样不需要再每次创建queueTemp~~
		for l > 0 {
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
			l--
		}
		// 遍历完当前层, 将结果集合并到result中
		result = append(result, list)
	}
	return result
}
