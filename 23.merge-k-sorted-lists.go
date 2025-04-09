/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30204
 *
 * [23] 合并 K 个升序链表
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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	// 使用分治法
	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return mergeTwoLists(left, right)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return dummy.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
    h := &ListNodeHeap{}
    heap.Init(h)

    // 初始化时将每个链表的头节点加入堆
    for _, node := range lists {
        if node != nil {
            heap.Push(h, node)
        }
    }

    dummy := &ListNode{}
    tail := dummy

    // 不断从堆中取出最小节点，接入结果链表
    for h.Len() > 0 {
        node := heap.Pop(h).(*ListNode)
        tail.Next = node
        tail = tail.Next
        if node.Next != nil {
            heap.Push(h, node.Next)
        }
    }

    return dummy.Next
}

// 定义一个最小堆
type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) }

func (h ListNodeHeap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
    old := *h
    n := len(old)
    node := old[n-1]
    *h = old[0 : n-1]
    return node
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/

