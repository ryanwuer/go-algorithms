/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=30204
 *
 * [300] 最长递增子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLIS(nums []int) int {
	f := make([]int, len(nums))
    for i, x := range nums {
        for j, y := range nums[:i] {
            if y < x {
                f[i] = max(f[i], f[j])
            }
        }
        f[i]++
    }
    return maxSlice(f)
}

func maxSlice(s []int) int {
	if len(s) == 0 {
		panic("slice is empty")
	}
	maxVal := s[0]
	for _, v := range s[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// @lc code=end

/*
// @lcpr case=start
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/

