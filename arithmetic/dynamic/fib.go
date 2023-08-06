package dynamic

// F(0) = 0，F(1) = 1
// F(n) = F(n - 1) + F(n - 2)，其中 n > 1
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// 加入了备忘录，时间复杂度变为 O(n)
func fib1(n int) int {
	var dp func(n int, memo []int) int
	dp = func(n int, memo []int) int {
		if n == 1 || n == 0 {
			return n
		}

		if memo[n] != 0 {
			return memo[n]
		}

		memo[n] = dp(n-1, memo) + dp(n-2, memo)
		return memo[n]
	}

	memo := make([]int, n+1)
	return dp(n, memo)
}

// 备忘录 + 迭代
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 压缩备忘录
func fib3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp0, dp1 := 0, 1
	var dpn int
	for i := 2; i <= n; i++ {
		dpn = dp0 + dp1

		dp0 = dp1
		dp1 = dpn
	}
	return dpn
}
