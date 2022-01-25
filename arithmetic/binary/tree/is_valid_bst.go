package tree

import "math"

// IsValidBST 验证二叉搜索树
// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
//	有效 二叉搜索树定义如下：
//		节点的左子树只包含 小于 当前节点的数。
//		节点的右子树只包含 大于 当前节点的数。
//		所有左子树和右子树自身必须也是二叉搜索树。
//
// ** 中序遍历得到一个递增的序列 **
//
//	有三种做法，
//		一种是使用全局变量，存储上一次访问到的节点，并使用递归。
//		一种是用全局变量，并使用迭代。
//		另外一种是使用类似动态规划的方式，在递归中将结果不断返回上层并修正的过程。这也是个二叉树解题模板。
// https://leetcode-cn.com/problems/validate-binary-search-tree
func IsValidBST(root *Node) bool {
	// 本来可以用全局变量来控制 preVal，但是 lc 上对全局变量处理的有问题
	// 所以只能改进，将preVal作为一个参数和返回值，在调用栈中传递
	isBst, _ := isValidBST(root, math.MinInt64)
	return isBst
}

func isValidBST(root *Node, preVal int) (bool, int) {
	if root == nil {
		// 这里应该返回 preVal，不能是 math.MinInt64
		// 因为对于 Nil 节点，不能覆盖之前访问的 preVal
		return true, preVal
	}

	// 访问左节点，并更新 preVal
	left, preVal := isValidBST(root.Left, preVal)

	// 判断条件
	if !left {
		return false, math.MinInt64
	}
	if preVal != math.MinInt64 && preVal >= root.Val {
		return false, math.MinInt64
	}

	// 更新 preVal
	preVal = root.Val

	// 访问右节点，并更新 preVal
	return isValidBST(root.Right, preVal)
}

// IsValidBSTII 判断搜索二叉树
// 使用迭代的方式
// 借助栈和额外变量
func IsValidBSTII(root *Node) bool {
	if root == nil {
		return true
	}

	// 下面是中序遍历迭代的模板
	stack := make([]*Node, 0)
	preVal := math.MinInt64

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 访问 current node，这里可以比较 preVal
			if preVal >= cur.Val {
				return false
			}
			preVal = cur.Val

			if cur.Right != nil {
				root = cur.Right
				continue
			}
		}
	}
	return true
}

// IsValidBSTIII 判断搜索二叉树
// 使用深度DP的方式，在递归过程中，将需要的结果返回，并不断优化结果
func IsValidBSTIII(root *Node) bool {
	result := isValidBSTIII(root)
	return result.IsValid
}

type IsValidBSTResult struct {
	IsValid bool
	Min     int
	Max     int
}

func isValidBSTIII(root *Node) *IsValidBSTResult {
	if root == nil {
		return nil
	}

	left := isValidBSTIII(root.Left)
	right := isValidBSTIII(root.Right)

	if left != nil && (!left.IsValid || left.Max >= root.Val) {
		return &IsValidBSTResult{false, 0, 0}
	}
	if right != nil && (!right.IsValid || right.Min <= root.Val) {
		return &IsValidBSTResult{false, 0, 0}
	}

	result := new(IsValidBSTResult)
	result.IsValid = true
	result.Min = root.Val
	result.Max = root.Val

	if left != nil && left.Min < result.Min {
		result.Min = left.Min
	}
	if right != nil && right.Max > result.Max {
		result.Max = right.Max
	}
	return result
}
