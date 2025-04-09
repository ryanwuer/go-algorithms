package main

import (
	"fmt"
)

/*
按格式合并两个链表

给定两个链表L1=a1→a2→⋯→an−1→an 和L2=b1→b2→⋯→bm−1→bm，其中n≥2m。

需要将较短的链表L2反转并合并到较长的链表L1中

使得合并后的链表如下形式：a1→a2→bm→a3→a4→bm−1→…

输入格式:
第一行包含两个链表的第一个节点地址（不确定哪个链表更长），以及两个链表的总节点数N(≤100000)。

节点地址用一个 5 位非负整数表示（可能有前导 0），空地址 NULL 用 −1 表示。

例如：00100 01000 7。其中00100表示第一个链表的首个节点地址，01000表示第二个链表的首个节点地址，7表示两个链表的总节点数。

接下来N行，每行描述一个节点的信息，格式如下：

Address Data Next

其中 Address 是节点地址，Data 是一个绝对值不超过100000的整数，Next 是下一个节点的地址。

保证两个链表都不为空，且较长的链表至少是较短链表长度的两倍。

输出格式:
对于每个测试用例，按顺序输出合并后的结果链表。每个结点占一行，按输入的格式输出。
*/

type Node struct {
	Address string
	Data    int
	Next    string
}

func main() {
	var head1, head2 string
	var totalNodes int
	fmt.Scan(&head1, &head2, &totalNodes)

	nodeMap := make(map[string]*Node)

	// Step 1: Parse input nodes
	for i := 0; i < totalNodes; i++ {
		var address, next string
		var data int
		fmt.Scan(&address, &data, &next)
		nodeMap[address] = &Node{Address: address, Data: data, Next: next}
	}

	// Step 2: Traverse both linked lists and store them in slices
	list1 := traverseList(head1, nodeMap)
	list2 := traverseList(head2, nodeMap)

	// Step 3: Ensure list1 is the longer one
	if len(list1) < len(list2) {
		list1, list2 = list2, list1
	}

	// Step 4: Reverse the shorter list (list2)
	reverseList(list2)

	// Step 5: Merge list2 into list1 according to the rules
	mergedList := mergeLists(list1, list2)

	// Step 6: Output the result
	for i := 0; i < len(mergedList)-1; i++ {
		fmt.Printf("%s %d %s\n", mergedList[i].Address, mergedList[i].Data, mergedList[i+1].Address)
	}
	// The last node's Next is -1 (NULL)
	fmt.Printf("%s %d -1\n", mergedList[len(mergedList)-1].Address, mergedList[len(mergedList)-1].Data)
}

// Traverse a linked list and return a slice of Node pointers
func traverseList(head string, nodeMap map[string]*Node) []*Node {
	var res []*Node

	for head != "-1" {
		n := nodeMap[head]
		res = append(res, n)
		head = n.Next
	}

	return res
}

// Reverse a slice of Node pointers
func reverseList(list []*Node) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

// Merge two lists according to the rules
func mergeLists(longList, shortList []*Node) []*Node {
	longLen := len(longList)
	shortLen := len(shortList)

	var res []*Node
	i, j := 0, 0
	for i < longLen {
		if i < longLen {
			res = append(res, longList[i])
			i++
		}
		if i < longLen {
			res = append(res, longList[i])
			i++
		}
		if j < shortLen {
			res = append(res, shortList[j])
			j++
		}
	}
	return res
}
