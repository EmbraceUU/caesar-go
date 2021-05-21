package linked_list

// DeleteDuplicates 给定一个按升序排序的链表
// 删除链表中重复的元素
// 错误记录：
// 1. 没有加上for循环
// 2. 没有加入当前节点p向后移动的逻辑，导致无限循环
// 3. 需要等重复节点全部删除完，再移到下一个节点
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func DeleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil && p.Next != nil {
		if p.Val != p.Next.Val {
			// 前后值不一样时，往后移动
			p = p.Next
			continue
		}
		// 前后值一样时，删除后面的值，不移动p
		pn := p.Next.Next
		p.Next.Next = nil
		p.Next = pn
	}
	return head
}

// DeleteDuplicatesLatest DeleteDuplicates最终版
func DeleteDuplicatesLatest(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

// DeleteDuplicatesII 给定一个按升序排序的链表，删除重复节点
// 1. 借助于dummy node来处理head
// 2. 借助于rmVal来删除所有重复节点
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func DeleteDuplicatesII(head *ListNode) *ListNode {
	// 使用dummy node来处理head
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	current := dummy
	// 链表里出现重复节点，最少有两个节点
	var rmVal int
	for current.Next != nil && current.Next.Next != nil {
		// 使用rmVal辅助变量来删除所有重复节点
		if current.Next.Val == current.Next.Next.Val {
			rmVal = current.Next.Val
			for current.Next != nil && current.Next.Val == rmVal {
				current.Next = current.Next.Next
			}
			continue
		}
		current = current.Next
	}
	return dummy.Next
}

// ReverseList 反转链表
// 使用一个fast和slow遍历list，并反转list
// 错误1. 没有加上slow.Next = nil，导致末尾出现无限循环
// https://leetcode-cn.com/problems/reverse-linked-list/
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head.Next
	slow := head
	slow.Next = nil
	for fast != nil {
		tmp := fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}

	return slow
}

// ReverseListLatest 反转链表，递归做法
// 关键是：递归到最后一个节点，作为新的头结点返回， 然后每一步都自己反转自己 head.Next.Next = head, 这样不需要借助其他节点。
func ReverseListLatest(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseListLatest(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// ReverseBetween 反转链表
// 反转left到right之间的节点， 包括left和right
// 错误：
// 1. 没有利用dummy node来统一处理head变化的情况
// 2. 思路不清晰，先遍历到left， 然后反转left到right，最后拼接left之前和right之后的节点
// https://leetcode-cn.com/problems/reverse-linked-list-ii/
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	depth := 1
	before := &ListNode{}
	after := &ListNode{}
	pre := head
	current := head.Next
	for current != nil && depth <= right {
		// 想到了遍历到left的步骤，但是应该提到循环外面去，并且需要借助dummy node，这样才能统一处理
		if depth < left {
			pre = current
			current = current.Next
			depth++
			continue
		}

		// 赋值before和after的部分思路不清晰
		// 在进入循环之前以及循环退出时处理即可
		if depth == left-1 {
			before = current
		} else if depth == right {
			after = current.Next
		}

		// 逻辑太耦合，需要兼顾before和after的情况
		// 对于before和after借助指针在外面处理即可
		if depth >= left {
			tmp := current.Next
			if depth == left {
				current.Next = after
			} else {
				current.Next = pre
			}
			if depth == right {
				before.Next = current
			} else {
				pre = current
			}
			current = tmp
		}

		depth++
	}

	return head
}

// ReverseBetweenLatest 反转链表
// 先遍历到 left 处，翻转，再拼接后续
func ReverseBetweenLatest(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	// 0->1->2->3->4->5->Nil
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy

	// 0->1(pre)->2(head)->3->4->5->Nil
	var pre *ListNode
	var i int
	for i < left {
		pre = head
		head = head.Next
		i++
	}

	j := i
	mid := head
	var next *ListNode
	for head != nil && j <= right {
		// 0->1 nil<-2(next) 3(head)->4->5->Nil j=2
		// 0->1 nil<-2<-3(next) 4(head)->5->Nil j=3
		// 0->1 nil<-2<-3<-4(next) 5(head)->Nil j=4
		tmp := head.Next
		head.Next = next
		next = head
		head = tmp
		j++
	}

	pre.Next = next
	mid.Next = head

	return dummy.Next
}

// MergeTwoLists 将两个升序链表合并为一个新的 升序 链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。
// 借助dummy node
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	// 同时遍历两个链表，从小到大拼接，直到有一个链表为nil
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	// 将另外一个不为nil的链表拼接在后面
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}

	return dummy.Next
}

// Partition 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
// 将大于等于x的节点放到另外一个链表上，然后连接两个节点
// 错误：没有处理好新链表的新增过程
// 错误：没有设置新链表最后一个节点为Nil, 导致了循环
// https://leetcode-cn.com/problems/partition-list/
func Partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var headr *ListNode
	dummyl := &ListNode{Next: head}
	dummyr := &ListNode{}
	head = dummyl
	headr = dummyr

	// 0(head)->1->4->3->2->5->2   0(headr)->Nil
	// 0->1(head)->4->3->2->5->2   0(headr)->Nil
	// 0->1->3(head)->2->5->2   0->4(headr)->Nil
	// 0->1->2(head)->5->2   0->4->3(headr)->Nil
	// 0->1->2->2(head)   0->4->3->5(headr)->Nil
	for head.Next != nil {
		if head.Next.Val >= x {
			headr.Next = head.Next
			headr = headr.Next // 4->3

			head.Next = head.Next.Next // 1->3
		} else {
			head = head.Next
		}
	}

	head.Next = dummyr.Next
	headr.Next = nil

	return dummyl.Next
}

// SortList 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
// 在 O(n log n) 时间复杂度和常数级空间复杂度下处理
// 思路：归并排序，将链表拆分，分到不能再分为止，然后合并
// https://leetcode-cn.com/problems/sort-list/
func SortList(head *ListNode) *ListNode {
	// -1->5->3->4->0
	// -1->5  3->4->0
	// -1->5  3  4->0
	// -1->5  3  0->4
	// -1->5  0->3->4
	// -1->0->3->4->5

	// 使用快慢指针找链表中点
	findMid := func(head *ListNode) *ListNode {
		// todo 之前用的head做fast，导致两个节点时陷入死循环
		fast := head.Next
		slow := head
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		return slow
	}

	// todo 因为没有处理mid与head的切分，导致head在mergeSort里陷入死循环
	var mergeSort func(node *ListNode) *ListNode
	mergeSort = func(head *ListNode) *ListNode {
		// 当只有一个节点时， 退出递归
		if head == nil || head.Next == nil {
			return head
		}

		mid := findMid(head)

		tail := mid.Next
		mid.Next = nil
		// 递归，得到排好序并且合并了的链表
		sortedHead := mergeSort(head)
		sortedMid := mergeSort(tail)

		// 合并链表
		dummy := &ListNode{}
		current := dummy
		for sortedHead != nil && sortedMid != nil {
			if sortedHead.Val <= sortedMid.Val {
				current.Next = sortedHead
				current = current.Next
				sortedHead = sortedHead.Next
			} else {
				current.Next = sortedMid
				current = current.Next
				sortedMid = sortedMid.Next
			}
		}

		if sortedHead != nil {
			current.Next = sortedHead
		}
		if sortedMid != nil {
			current.Next = sortedMid
		}

		return dummy.Next
	}
	return mergeSort(head)
}

// ReorderList 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
// 思路：
// 1. 找到中点，将链表切成两段
// 2. 将后半段反转
// 3. 合并两个链表
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 1->2->3->4->5
	// 1->2->3(slow)->4->5->nil(fast)

	// 1->2->3 4(mid)->5->nil(fast)
	mid := slow.Next
	slow.Next = nil

	// nil(next)  4(cur)->5->nil
	// nil<-4(next) 5(cur)->nil
	// nil<-4<-5(next) nil(cur)
	current := mid
	var next *ListNode
	for current != nil {
		tmp := current.Next
		current.Next = next
		next = current
		current = tmp
	}

	// 1->5->2(cur)->4(next)->3
	current = head
	for next != nil {
		tmp := current.Next // nil
		tmp2 := next.Next   // nil
		current.Next = next
		next.Next = tmp
		next = tmp2
		current = tmp
	}
}

// HasCycle 给定一个链表，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
// https://leetcode-cn.com/problems/linked-list-cycle/
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

// DetectCycle 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 错误：忽略了值相同的情况
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
func DetectCycle(head *ListNode) *ListNode {
	// [-1,-7,7,-4,19,6,-9(cy),-5,-2,-5]
	// 6
	if head == nil || head.Next == nil {
		return nil
	}
	nodeMap := make(map[*ListNode]struct{})
	for head != nil {
		_, ok := nodeMap[head]
		if ok {
			return head
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}
