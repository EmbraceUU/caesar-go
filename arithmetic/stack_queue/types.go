package stack_queue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val       int
	Neighbors []*Node
}
