/*
 * @lc app=leetcode.cn id=31 lang=golang
 * @lcpr version=30204
 *
 * [31] 下一个排列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func nextPermutation(nums []int) {
	// Find the first decreasing element from the end
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// If we found a decreasing element
	if i >= 0 {
		// Find the next larger element to swap with
		j := len(nums) - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		// Swap the two elements
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Reverse the elements after the swapped position
	reverse(nums, i+1)
}

// Helper function to reverse elements in the slice
func reverse(nums []int, start int) {
	for i, j := start, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,5]\n
// @lcpr case=end

*/

