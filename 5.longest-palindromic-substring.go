/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30204
 *
 * [5] 最长回文子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) string {
	// 处理特殊情况
	if len(s) < 2 {	
		return s
	}
	// 初始化dp数组
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	// 初始化最长回文子串的起始位置和长度
	start, maxLen := 0, 1
	// 遍历字符串
	for j := 0; j < len(s); j++ {	
		for i := j; i >= 0; i-- {
			// 判断是否是回文子串
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				// 更新最长回文子串的起始位置和长度
				if j-i+1 > maxLen {
					start = i
					maxLen = j - i + 1
				}
			}
		}
	}
	// 返回最长回文子串
	return s[start : start+maxLen]
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

