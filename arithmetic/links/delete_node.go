package links

// deleteNode 请编写一个函数，用于 删除单链表中某个特定节点 。在设计函数时需要注意，你无法访问链表的头节点head ，只能直接访问 要被删除的节点 。
// 题目数据保证需要删除的节点 不是末尾节点 。
// 思路：改头换面
// https://leetcode-cn.com/problems/delete-node-in-a-linked-list
func deleteNode(node *ListNode) {
	//	1 -> 2 -> 3 -> 4 -> 5 -> 6
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
