/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=30204
 *
 * [39] 组合总和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(start int, path []int, sum int)
	dfs = func(start int, path []int, sum int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				continue
			}
			path = append(path, candidates[i])
			dfs(i, path, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/

