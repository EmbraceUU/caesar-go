package links

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	//为什么我们要规定初始时慢指针在位置 head，快指针在位置 head.next，而不是两个指针都在位置 head（即与「乌龟」和「兔子」中的叙述相同）？
	//
	//观察下面的代码，我们使用的是 while 循环，循环条件先于循环体。
	//由于循环条件一定是判断快慢指针是否重合，如果我们将两个指针初始都置于 head，那么 while 循环就不会执行。
	//因此，我们可以假想一个在 head 之前的虚拟节点，慢指针从虚拟节点移动一步到达 head，快指针从虚拟节点移动两步到达 head.next，
	//这样我们就可以使用 while 循环了。
	fast, slow := head.Next, head
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
