package tree

// Node TreeNode 节点类型
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewTreeNode(nums []int) *Node {
	var root *Node

	l := len(nums)
	if l == 0 {
		return root
	}

	root = &Node{Val: nums[0]}
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for i := 0; i < l; i += 2 {
		p := queue[0]
		queue = queue[1:]

		// 添加左节点
		if i+1 < l && nums[i+1] > 0 {
			p.Left = &Node{Val: nums[i+1]}
			queue = append(queue, p.Left)
		}

		// 添加右节点
		if i+2 < l && nums[i+2] > 0 {
			p.Right = &Node{Val: nums[i+2]}
			queue = append(queue, p.Right)
		}
	}
	return root
}
