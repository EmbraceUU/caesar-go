package tree

// IsCompleteTree 判断完全二叉树
// 若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，
// 第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~2h个节点。）
// 两个条件：
//  1）右节点存在时，左节点必存在
//  2）当出现叶子节点后，以后必都是叶子节点
// https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
func IsCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// [1,  2,3,   5,null, 7,8   ]

	// 利用队列，使用广度优先遍历

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	leaf := false

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// 将原本访问节点的地方，改为判断条件

		// 如果右节点存在，左节点不存在，那么不符合
		if cur.Right != nil && cur.Left == nil {
			return false
		}
		// 如果已经是叶子节点，还存在子节点，不符合
		if leaf && (cur.Left != nil || cur.Right != nil) {
			return false
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

		// 当右节点为nil时，标记叶子节点开始
		if cur.Right == nil {
			leaf = true
		}
	}
	return true
}
