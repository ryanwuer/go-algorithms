/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=30204
 *
 * [121] 买卖股票的最佳时机
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

