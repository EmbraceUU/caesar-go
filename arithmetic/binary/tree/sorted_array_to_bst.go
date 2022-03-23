package tree

// sortedArrayToBST 108. 将有序数组转换为二叉搜索树
// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
func sortedArrayToBST(nums []int) *Node {
	// 分治法，二分法，将 nums 分为两部分，mid 为 root 节点
	// 左右两侧再二分，left mid 为 left ，right mid 为 right
	// 再细分下去
	var divide func(start, end int, nums []int) *Node
	divide = func(start, end int, nums []int) *Node {
		mid := start + (end-start)>>1
		root := new(Node)
		root.Val = nums[mid]

		if mid-1 >= start {
			root.Left = divide(start, mid-1, nums)
		}
		if mid+1 <= end {
			root.Right = divide(mid+1, end, nums)
		}
		return root
	}

	return divide(0, len(nums)-1, nums)
}
