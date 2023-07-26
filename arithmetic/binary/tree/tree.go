package tree

// LevelOrderBottom 自下而上BFS
// 使用LeverOrder的方法，把结果反转即可
func LevelOrderBottom(root *TreeNode) [][]int {
	result := LevelOrderII(root)
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
func ZigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var toggle bool

	for len(queue) > 0 {
		list := make([]int, 0)
		queueTemp := make([]*TreeNode, 0)

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

func MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}

// --------------------------- 分界线 --------------------------- //

// MaxPathSum 二叉树中的最大路径和
// 路径每到一个节点，有 3 种选择：1. 停在当前节点。2. 走到左子节点。3. 走到右子节点。
// 使用dfs, 每次更新maxSum，然后每次返回该节点的最大路径，并且负数时返回0
func MaxPathSum(root *TreeNode) int {
	var maxSum int
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
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

func IsSymmetric(root *TreeNode) bool {
	return isSymmetricIterative(root)
}

func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST(nums)
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
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := make([]*TreeNode, 0)
	qPath := make([]*TreeNode, 0)

	var findPath func(root, p *TreeNode, path []*TreeNode) []*TreeNode

	findPath = func(root, p *TreeNode, path []*TreeNode) []*TreeNode {
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
func LowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
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

// InsertIntoBST 插入二叉搜索树
// 将新插入的节点插入到叶子节点即可
func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	node := root
	for node != nil {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{
					Val: val,
				}
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{
					Val: val,
				}
				break
			}
			node = node.Right
		}
	}
	return root
}
