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

func TestQuickSort(t *testing.T) {
	nums1 := [...]int{4, 10, 2, 5, 9, 6, 7, 3, 8, 1}
	nums := make([]int, 0)
	for _, v := range nums1 {
		nums = append(nums, v)
	}
	nums = QuickSort(nums)
	fmt.Println(nums)
}

func TestSelectionSort(t *testing.T) {
	nums1 := [...]int{4, 10, 2, 5, 9, 6, 7, 3, 8, 1}
	nums := make([]int, 0)
	for _, v := range nums1 {
		nums = append(nums, v)
	}
	SelectionSort(nums)
	fmt.Println(nums)
}

func TestBubbleSort(t *testing.T) {
	nums1 := [...]int{4, 10, 2, 5, 9, 6, 7, 3, 8, 1}
	nums := make([]int, 0)
	for _, v := range nums1 {
		nums = append(nums, v)
	}
	BubbleSort(nums)
	fmt.Println(nums)
}

func TestInsertionSort(t *testing.T) {
	nums := []int{4, 10, 2, 5, 9, 6, 7, 3, 8, 1}
	InsertionSort(nums)
	fmt.Println(nums)
}

func TestFindPeakElement(t *testing.T) {
	nums1 := [...]int{1, 2}
	nums := make([]int, 0)
	for _, v := range nums1 {
		nums = append(nums, v)
	}
	fmt.Println(FindPeakElement(nums))
}

func TestQuickSortII(t *testing.T) {
	//nums := []int{3, 5, 1, 4, 29, 12, 2, 6}
	nums := []int{13, 5}
	QuickSortII(nums)
	t.Log(nums)
}

func TestHeapSort(t *testing.T) {
	nums := []int{3, 5, 1, 4, 29, 12, 2, 6}
	//nums := []int{13, 5}
	HeapSort(nums)
	t.Log(nums)
}

func TestContainsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 0, 1, 1}
	k := 1
	t.Log(ContainsNearbyDuplicate(nums, k))
}
