/*
 * @lc app=leetcode.cn id=29 lang=golang
 * @lcpr version=30204
 *
 * [29] 两数相除
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	dividend, divisor = abs(dividend), abs(divisor)
	result := 0

	for dividend >= divisor {
		dividend -= divisor
		result++
	}

	return result * sign
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// 10\n3\n
// @lcpr case=end

// @lcpr case=start
// 7\n-3\n
// @lcpr case=end

*/

