/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=30204
 *
 * [53] 最大子数组和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSubArray(nums []int) int {
	// dp[i]表示以nums[i]结尾的最大子数组和
	// dp[i] = max(dp[i-1]+nums[i], nums[i])
	// ans = max(ans, dp[i])
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/

