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
	// 这里开启的一个协程, 一直往ch里面写递增的自然数
	go generate(ch)
	for {
		// 从当前ch中取出第一个作为质子
		prime := <-ch
		// 打印当前质子
		fmt.Println(prime)
		// 新创建一个chan ch1
		ch1 := make(chan int)
		// 开启一个协程, 将当前的ch中的数值按照prime质子过滤, 写入到ch1中
		go filter(ch, ch1, prime)
		// 将ch1的引用给ch, 继续下一个质子的循环过滤, 并在下一个循环中, 新创建一个ch1, 将当前ch1的数据进行过滤, 这样类推...
		ch = ch1
		// 所以这里ch的指针是一直在变的, 但是generate中的ch的指针一直没有变, 因为它作为第一个ch, 一直在往里面写数据, 然后传递给后面的ch1,
		fmt.Printf("ch01: %p \n", ch)
	}
}

/*
以2和3作为质子的例子
*/
func Sieve01() {
	// 创建一个ch
	ch := make(chan int)
	// 开启一个协程, 一直往ch中写自然数
	go generate(ch)

	// 从ch中读取一个数, 作为质子, 这里是2
	prime := <-ch
	// 打印2
	fmt.Println(prime)
	// 新创建一个chan ch1
	ch1 := make(chan int)
	// 开启一个协程, 将当前的ch的数值, 按照2过滤, 写入到ch1中
	go filter(ch, ch1, prime)
	// 将ch1的应用指给ch
	ch = ch1

	// 从ch1中读出数值, 作为质子, 这里是3
	prime = <-ch
	fmt.Println(prime)
	// 新创建一个chan 指给ch1
	ch1 = make(chan int)
	// 将ch(原ch1) 在新协程中按照3过滤, 写入到新ch1中
	go filter(ch, ch1, prime)

	for i := 0; i < 10; i++ {
		p := <-ch1
		fmt.Println(p)
	}
}
