/*
 * @lc app=leetcode.cn id=16 lang=golang
 * @lcpr version=30204
 *
 * [16] 最接近的三数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}
	}

	return closest
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
// [-1,2,1,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n1\n
// @lcpr case=end

*/

