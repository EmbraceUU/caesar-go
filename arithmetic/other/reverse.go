package other

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
