package links

import (
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
