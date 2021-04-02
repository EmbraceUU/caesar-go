package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums1 := [...]int{2, 4, 1, 5, 3, 8, 6}
	nums := make([]int, 0)
	for _, v := range nums1 {
		nums = append(nums, v)
	}
	fmt.Println(MergeSort(nums))
}
