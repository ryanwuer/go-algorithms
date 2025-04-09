/*
 * @lc app=leetcode.cn id=24 lang=golang
 * @lcpr version=30204
 *
 * [24] 两两交换链表中的节点
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next

		// 交换节点
		first.Next = second.Next
		second.Next = first
		prev.Next = second

		prev = first
	}

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

