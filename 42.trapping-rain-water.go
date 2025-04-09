/*
 * @lc app=leetcode.cn id=42 lang=golang
 * @lcpr version=30204
 *
 * [42] 接雨水
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	trappedWater := 0
	for i := 0; i < n; i++ {
		trappedWater += min(leftMax[i], rightMax[i]) - height[i]
	}

	return trappedWater
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,2,1,0,1,3,2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,0,3,2,5]\n
// @lcpr case=end

*/

