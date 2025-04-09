package main

/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=30204
 *
 * [46] 全排列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

// Solution: Using used-element tracking backtracking
func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))

	// backtrack is a recursive helper function that builds permutations
	// by selecting unused elements one by one
	var backtrack func()
	backtrack = func() {
		// Base case: when we've used all elements
		if len(path) == len(nums) {
			// Create a copy of current permutation and add to result
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// Try all possible elements that haven't been used yet
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				// Mark current element as used
				used[i] = true
				// Add current element to current path
				path = append(path, nums[i])
				// Recursively build the rest of the permutation
				backtrack()
				// Backtrack: remove current element and mark as unused
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}

	// Start building permutations
	backtrack()
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
