package array

// maxProfit 买卖股票的最佳时机 II
// 【数组】【动态规划】
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	// 最终的结果，必然是最后一天将手里的股票卖出利益最大
	// 中间的每一天，利益取决于 是卖出还是买入，
	var dp0, dp1 = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		// dp0 = max(dp0, dp1+prices[i])
		// dp1 = max(dp1, dp0-prices[i])
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}

// maxProfitII 买卖股票的最佳时机 II
// 【数组】【贪心】
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfitII(prices []int) int {
	// 相当于寻找 x 个不想交的区间，使得利益的综合最大化
	// 而 li, ri区间的利益，等价于(li ,li+1],(li+1,li+2],…,(ri−1,ri]
	var ans int
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
