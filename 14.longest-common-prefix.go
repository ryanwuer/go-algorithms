package main

/*
 * @lc app=leetcode.cn id=14 lang=golang
 * @lcpr version=30203
 *
 * [14] 最长公共前缀
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		var curStr = strs[i]
		for j := 0; j < len(prefix) && j < len(curStr); j++ {
			if prefix[j] != curStr[j] {
				prefix = prefix[:j]
				break
			}
		}
		if len(prefix) > len(curStr) {
			prefix = prefix[:len(curStr)]
		}
	}
	return prefix
}

// @lc code=end

/*
// @lcpr case=start
// ["flower","flow","flight"]\n
// @lcpr case=end

// @lcpr case=start
// ["dog","racecar","car"]\n
// @lcpr case=end

*/

