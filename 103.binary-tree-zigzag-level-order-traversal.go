/*
 * @lc app=leetcode.cn id=103 lang=golang
 * @lcpr version=30204
 *
 * [103] 二叉树的锯齿形层序遍历
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if leftToRight {
				level[i] = node.Val
			} else {
				level[levelSize-i-1] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
		leftToRight = !leftToRight
	}

	return res
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
    return dfs(root, 0, [][]int{})
}

func dfs(node *TreeNode, depth int, res [][]int) [][]int {
    if node == nil {
        return res
    }

    // 如果当前层不存在，就创建一个
    if len(res) <= depth {
        res = append(res, []int{})
    }

    if depth%2 == 0 {
        res[depth] = append(res[depth], node.Val)
    } else {
        // 手动插入到头部
        level := res[depth]
        level = append([]int{node.Val}, level...)
        res[depth] = level
    }

    res = dfs(node.Left, depth+1, res)
    res = dfs(node.Right, depth+1, res)

    return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

