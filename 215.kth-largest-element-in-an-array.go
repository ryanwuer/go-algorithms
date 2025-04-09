/*
 * @lc app=leetcode.cn id=215 lang=golang
 * @lcpr version=30204
 *
 * [215] 数组中的第K个最大元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/

