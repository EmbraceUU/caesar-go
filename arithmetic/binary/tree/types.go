package tree

// TreeNode 节点类型
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(nums []int) *TreeNode {
	var root *TreeNode

	l := len(nums)
	if l == 0 {
		return root
	}

	root = &TreeNode{Val: nums[0]}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for i := 0; i < l; i += 2 {
		p := queue[0]
		queue = queue[1:]

		// 添加左节点
		if i+1 < l && nums[i+1] > 0 {
			p.Left = &TreeNode{Val: nums[i+1]}
			queue = append(queue, p.Left)
		}

		// 添加右节点
		if i+2 < l && nums[i+2] > 0 {
			p.Right = &TreeNode{Val: nums[i+2]}
			queue = append(queue, p.Right)
		}
	}
	return root
}
