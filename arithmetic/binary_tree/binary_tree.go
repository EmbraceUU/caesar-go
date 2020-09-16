package binary_tree

import "fmt"

// 前序递归
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)

	// 遍历左子树
	preorderTraversal(root.Left)
	// 遍历右子树
	preorderTraversal(root.Right)
}

// 前序非递归遍历
func preorderTraversal2(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	// 先根节点, 再左节点 再右节点
	for root != nil && root.Left != nil {
		fmt.Println(root.Left.Val)
		root = root.Left
	}

}
