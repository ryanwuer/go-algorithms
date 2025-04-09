/*
 * @lc app=leetcode.cn id=LCR 103 lang=golang
 * @lcpr version=30204
 *
 * [LCR 103] 零钱兑换
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化 dp 数组
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	// dp[0] = 0
	dp[0] = 0
	// dp[i] = min(dp[i], dp[i-coin]+1)
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// @lc code=end

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1]\n2\n
// @lcpr case=end

*/

