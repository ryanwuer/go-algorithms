package main

/*
 * @lc app=leetcode.cn id=9 lang=golang
 * @lcpr version=30203
 *
 * [9] 回文数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var y int
	for x > y {
		y = y*10 + x%10
		x = x/10
	}
	return x == y || x == y/10
}

// @lc code=end

/*
// @lcpr case=start
// 121\n
// @lcpr case=end

// @lcpr case=start
// -121\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

*/

