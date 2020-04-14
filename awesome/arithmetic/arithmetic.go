package arithmetic

import "fmt"

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}

func Reverse(x int) int {
	const IntMax = int(^uint(0)>>32) - 1
	const IntMin = ^IntMax
	var rev int
	for x != 0 {
		pop := x % 10
		x = x / 10
		if rev > IntMax/10 || rev == IntMax/10 && pop > 7 {
			return 0
		}
		if rev < IntMin/10 || rev == IntMin/10 && pop < -8 {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

/*
素数筛选算法
go语言本身有并发的特性, 所以可以并发的去做
从1到n的自然数, 从2开始计算, 然后用2开始每个素数
*/
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // send i to channel ch
	}
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i // send i to channel dst
		}
	}
}

func Sieve() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
