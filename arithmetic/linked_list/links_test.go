package linked_list

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
	numsArr := [...]int{1, 2, 3, 4, 5}
	nums := make([]int, 0)
	for _, v := range numsArr {
		nums = append(nums, v)
	}

	head := NewListNode(nums)

	head = ReverseList(head)
	fmt.Println(head)
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
	numsArr := [...]int{}
	nums := make([]int, 0)
	for _, v := range numsArr {
		nums = append(nums, v)
	}

	l1 := NewListNode(nums)

	numsArr2 := [...]int{}
	nums = make([]int, 0)
	for _, v := range numsArr2 {
		nums = append(nums, v)
	}

	l2 := NewListNode(nums)

	head := MergeTwoLists(l1, l2)

	fmt.Println(head)
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
