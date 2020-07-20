package main

const (
	max     = 50000000
	block   = 500
	bufSize = 100
)

// 非数据块的方式, 一个一个的传输
func NonBlock() {
	done := make(chan struct{})
	c := make(chan int, bufSize)

	go func() {
		count := 0
		for x := range c {
			count += x
		}
		close(done)
	}()

	for i := 0; i < max; i++ {
		c <- i
	}

	close(c)
	<-done
}

// 数据块的方式, 一块一块的传输
func Block() {
	done := make(chan struct{})
	c := make(chan [block]int, bufSize)

	go func() {
		count := 0
		for a := range c {
			for x := range a {
				count += x
			}
		}
		close(done)
	}()

	for i := 0; i < max; i += block {
		var b [block]int
		for n := 0; n < block; n++ {
			b[n] = n + i
			if n+i == max-1 {
				break
			}
		}
		c <- b
	}

	close(c)
	<-done
}
