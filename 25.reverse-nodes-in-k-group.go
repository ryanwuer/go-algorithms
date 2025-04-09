/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=30204
 *
 * [25] K 个一组翻转链表
 */

package main

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

// reverseKGroup 将链表每 k 个节点一组进行翻转
// 如果节点总数不是 k 的整数倍，那么最后剩余的节点保持原有顺序
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 处理特殊情况：空链表或 k=1（不需要翻转）
	if head == nil || k == 1 {
		return head
	}

	// 创建一个虚拟头节点，用于处理从第一个节点开始翻转的情况
	dummy := &ListNode{Next: head}
	// 记录上一组的最后一个节点
	prevGroupEnd := dummy

	for {
		// 获取当前组的起始节点
		groupStart := prevGroupEnd.Next
		// 初始化当前组的结束节点为上一组的结束节点
		groupEnd := prevGroupEnd

		// 移动 groupEnd 到当前组的第 k 个节点
		for i := 0; i < k && groupEnd != nil; i++ {
			groupEnd = groupEnd.Next
		}

		// 如果找不到 k 个节点，说明剩余节点不足 k 个，直接退出
		if groupEnd == nil {
			break
		}

		// 保存下一组的起始节点
		nextGroupStart := groupEnd.Next
		// 断开当前组与下一组的连接，以便单独处理当前组
		groupEnd.Next = nil

		// 翻转当前组，并将翻转后的组连接到上一组的末尾
		prevGroupEnd.Next = reverseList(groupStart)
		// 将翻转后的组的末尾连接到下一组的起始节点
		groupStart.Next = nextGroupStart

		// 更新上一组的末尾节点为当前组的末尾（即翻转前的起始节点）
		prevGroupEnd = groupStart
	}

	return dummy.Next
}

// reverseList 翻转一个链表并返回新的头节点
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	// 遍历链表，逐个翻转节点的 Next 指针
	for current != nil {
		// 保存下一个节点，因为我们要修改当前节点的 Next 指针
		next := current.Next
		// 翻转当前节点的 Next 指针
		current.Next = prev
		// 移动 prev 和 current 指针
		prev = current
		current = next
	}

	// 返回新的头节点（原链表的最后一个节点）
	return prev
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end
*/
