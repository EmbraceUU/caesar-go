package tree

// LevelOrder 【BFS 广度优先搜索 第一版】
func LevelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		var level []int
		var queueTemp []*Node
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

// LevelOrderBottom 自下而上BFS
// 使用LeverOrder的方法，把结果反转即可
func LevelOrderBottom(root *Node) [][]int {
	result := LevelOrder(root)
	reverse(result)
	return result
}

// 反转列表
func reverse(result [][]int) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

// ZigzagLevelOrder 自上而下, Z字形遍历
func ZigzagLevelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	var toggle bool

	for len(queue) > 0 {
		list := make([]int, 0)
		queueTemp := make([]*Node, 0)

		for i := 0; i < len(queue); i++ {
			list = append(list, queue[i].Val)
			if queue[i].Left != nil {
				queueTemp = append(queueTemp, queue[i].Left)
			}
			if queue[i].Right != nil {
				queueTemp = append(queueTemp, queue[i].Right)
			}
		}

		if toggle {
			reverseII(list)
		}

		result = append(result, list)
		queue = queueTemp
		toggle = !toggle
	}

	return result
}

func reverseII(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// --------------------------- 分界线 --------------------------- //

// MaxDepth 求二叉树的最大深度
func MaxDepth(root *Node) int {
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
func IsBalanced(root *Node) bool {
	if root == nil {
		return true
	}

	depth := isBalanced(root)
	if depth < 0 {
		return false
	}
	return true
}

func isBalanced(root *Node) int {
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

// MaxPathSum 二叉树中的最大路径和
// 路径每到一个节点，有 3 种选择：1. 停在当前节点。2. 走到左子节点。3. 走到右子节点。
// 使用dfs, 每次更新maxSum，然后每次返回该节点的最大路径，并且负数时返回0
func MaxPathSum(root *Node) int {
	var maxSum int
	var dfs func(root *Node) int
	dfs = func(root *Node) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		maxSum = max(maxSum, left+root.Val+right)

		return max(0, root.Val+max(left, right))
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LowestCommonAncestor 我认为，就是找到一个共同的节点
// 应该是从下往上找,
// [3,5,1,6,2,0,8,null,null,7,4]
//	5
//	4
func LowestCommonAncestor(root, p, q *Node) *Node {
	pPath := make([]*Node, 0)
	qPath := make([]*Node, 0)

	var findPath func(root, p *Node, path []*Node) []*Node

	findPath = func(root, p *Node, path []*Node) []*Node {
		if root == nil {
			return path
		}
		if p == nil {
			return path
		}

		if root != p {
			// 这里没有重新赋值
			// 左子树和右子树 不会同时都有路径, 这个是关键
			pathl := findPath(root.Left, p, path)
			pathr := findPath(root.Right, p, path)

			if len(pathl) > 0 {
				path = pathl
			} else if len(pathr) > 0 {
				path = pathr
			}

			if len(path) > 0 {
				path = append(path, root)
			}
		} else {
			path = append(path, root)
			return path
		}

		return path
	}

	pPath = findPath(root, p, pPath)
	qPath = findPath(root, q, qPath)

	pLen, qLen := len(pPath), len(qPath)

	for i := 0; i < pLen; i++ {
		for j := 0; j < qLen; j++ {
			if pPath[i] == qPath[j] {
				return pPath[i]
			}
		}
	}
	return root
}

// LowestCommonAncestorII
// 通过递归对二叉树进行后序遍历，当遇到节点 p 或 q 时返回。从底至顶回溯，
// 当节点 p, q 在节点 root 的异侧时，节点 root 即为最近公共祖先，则向上返回 root。
func LowestCommonAncestorII(root, p, q *Node) *Node {
	// 如果root为nil，返回nil
	if root == nil {
		return nil
	}
	// 如果root为p， 那么root为p的最近的父节点; 对q亦然
	if root == p || root == q {
		return root
	}
	// 开启递归过程, 向左右子树寻找公共祖先
	left := LowestCommonAncestorII(root.Left, p, q)
	right := LowestCommonAncestorII(root.Right, p, q)

	// 如果p和q分别在root的两侧，节点root即为最近公共祖先
	if left != nil && right != nil {
		return root
	}

	// 再往上找，肯定会出现left或者right为Nil的情况
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}

	// 如果left和right都为nil, 则返回nil， 说明root下面没有p或者q
	return nil
}

// IsValidBST 验证二叉搜索树
// 思路: 左子树返回一个最大的数, 右子树返回一个最小的数, 然后和根节点比较
func IsValidBST(root *Node) bool {
	_, _, isBST := maxAndMin(root)
	return isBST
}

func maxAndMin(root *Node) (int, int, bool) {
	if root == nil {
		return 0, 0, true
	}

	lmax, lmin, lb := maxAndMin(root.Left)
	rmax, rmin, rb := maxAndMin(root.Right)

	// 如果子树有不满足的, 直接返回false
	if !lb || !rb {
		return 0, 0, false
	}

	if rmin != 0 && rmin <= root.Val {
		return 0, 0, false
	}

	if lmax != 0 && lmax >= root.Val {
		return 0, 0, false
	}

	if rmax == 0 {
		rmax = root.Val
	}
	if lmin == 0 {
		lmin = root.Val
	}

	return rmax, lmin, true
}

// IsValidBSTII 验证二叉搜索树
// 使用中序遍历二叉树，然后判断排序是否正确
func IsValidBSTII(root *Node) bool {
	result := make([]int, 0)
	inOrder(root, &result)

	l := len(result)
	for i := 0; i < l; i++ {
		if i < l-1 {
			if result[i] > result[i+1] {
				return false
			}
		}
	}
	return true
}

func inOrder(root *Node, result *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}

// InsertIntoBST 插入二叉搜索树
// 将新插入的节点插入到叶子节点即可
func InsertIntoBST(root *Node, val int) *Node {
	if root == nil {
		return &Node{
			Val: val,
		}
	}

	node := root
	for node != nil {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &Node{
					Val: val,
				}
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &Node{
					Val: val,
				}
				break
			}
			node = node.Right
		}
	}
	return root
}
