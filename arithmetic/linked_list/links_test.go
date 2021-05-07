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
