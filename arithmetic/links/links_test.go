package links

import (
	"fmt"
	"testing"
)

func TestNewListNode(t *testing.T) {
	numsArr := [...]int{1, 1, 1}
	nums := make([]int, 0)
	for _, v := range numsArr {
		nums = append(nums, v)
	}

	head := NewListNode(nums)
	head = DeleteDuplicates(head)
}

func TestDeleteDuplicatesII(t *testing.T) {
	numsArr := [...]int{1, 1, 1}
	nums := make([]int, 0)
	for _, v := range numsArr {
		nums = append(nums, v)
	}

	head := NewListNode(nums)

	head = DeleteDuplicatesII(head)

	fmt.Println(head)
}

func TestReverseList(t *testing.T) {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	head = ReverseList(head)
	t.Log(PrintListNode(head))
}

func TestReverseBetween(t *testing.T) {
	numsArr := [...]int{1, 2, 3, 4, 5}
	nums := make([]int, 0)
	for _, v := range numsArr {
		nums = append(nums, v)
	}

	head := NewListNode(nums)

	head = ReverseBetween(head, 2, 4)

	fmt.Println(head)
}

func TestMergeTwoLists(t *testing.T) {
	l1 := NewListNode([]int{1, 3, 5, 7, 9})
	l2 := NewListNode([]int{2, 4})
	head := MergeTwoLists(l1, l2)
	t.Log(PrintListNode(head))
}

func TestPartition(t *testing.T) {
	numsArr := [...]int{1, 4, 3, 2, 5, 2}
	nums := make([]int, 0)
	for _, v := range numsArr {
		nums = append(nums, v)
	}

	head := NewListNode(nums)

	head = Partition(head, 3)

	fmt.Println(head)
}

func TestSortList(t *testing.T) {
	numsArr := [...]int{-1, 5, 3, 4, 0}
	nums := make([]int, 0)
	for _, v := range numsArr {
		nums = append(nums, v)
	}

	head := NewListNode(nums)

	head = SortList(head)

	fmt.Println(head)
}

func TestReorderList(t *testing.T) {
	numsArr := [...]int{1, 2, 3}
	nums := make([]int, 0)
	for _, v := range numsArr {
		nums = append(nums, v)
	}

	head := NewListNode(nums)

	ReorderList(head)

	fmt.Println(head)
}

func TestIsPalindrome(t *testing.T) {
	head := NewListNode([]int{1, 0, 1, 2})
	is := IsPalindrome(head)
	fmt.Println(is)
}

func TestAddTwoNumbers(t *testing.T) {
	a1 := []int{9, 9, 9, 9, 9}
	a2 := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	l1 := NewListNode(a1)
	l2 := NewListNode(a2)

	t.Log(PrintListNode(AddTwoNumbers(l1, l2)))
}

func TestDeleteNode(t *testing.T) {
	l1 := NewListNode([]int{1, 2, 3, 4, 5, 6})
	DeleteNode(l1.Next.Next.Next)
	t.Log(PrintListNode(l1))
}

func TestRemoveNthFromEnd(t *testing.T) {
	l1 := NewListNode([]int{1, 2, 3, 4, 5, 6, 7})
	n := 7
	l1 = RemoveNthFromEnd(l1, n)
	t.Log(PrintListNode(l1))
}
