package main

/*
 * @lc app=leetcode.cn id=160 lang=golang
 * @lcpr version=30203
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// Initialize two pointers, one for each list
	pa, pb := headA, headB

	// Continue until both pointers meet (either at intersection or nil)
	for pa != pb {
		// If pa reaches end of list A, switch to list B
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}

		// If pb reaches end of list B, switch to list A
		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}

	// Return the intersection node (or nil if no intersection)
	return pa
}

// @lc code=end

/*
// @lcpr case=start
// 8\n[4,1,8,4,5]\n[5,6,1,8,4,5]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// 2\n[1,9,1,2,4]\n[3,2,4]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 0\n[2,6,4]\n[1,5]\n3\n2\n
// @lcpr case=end

*/
