/*
 * @lc app=leetcode.cn id=124 lang=golang
 * @lcpr version=30204
 *
 * [124] 二叉树中的最大路径和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	// 动态规划实现
	maxSum := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))
		maxSum = max(maxSum, left+right+node.Val)
		return max(left, right) + node.Val
	}
	dfs(root)
	return maxSum
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [-10,9,20,null,null,15,7]\n
// @lcpr case=end

*/

