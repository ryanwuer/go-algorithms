/*
 * @lc app=leetcode.cn id=518 lang=golang
 * @lcpr version=30204
 *
 * [518] 零钱兑换 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func change(amount int, coins []int) int {
	// dp[i]表示凑成金额i的组合数
	dp := make([]int, amount+1)
	// 0元的组合数为1
	dp[0] = 1

	// 遍历硬币
	for _, coin := range coins {
		// 遍历金额
		for j := coin; j <= amount; j++ {
			// 更新组合数
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

// @lc code=end

/*
// @lcpr case=start
// 5\n[1, 2, 5]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[2]\n
// @lcpr case=end

// @lcpr case=start
// 10\n[10]\n
// @lcpr case=end

*/

