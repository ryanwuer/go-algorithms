/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30204
 *
 * [2] 两数相加
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位
	carry := 0
	// 哑节点
	dummy := &ListNode{}
	// 当前节点
	current := dummy

	for l1 != nil || l2 != nil || carry > 0 {
		// 获取当前节点的值
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		// 计算当前节点的值和进位
		sum := val1 + val2 + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/

