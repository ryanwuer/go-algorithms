package main

import (
	"fmt"
)

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
	var list []*Node
	for head != "-1" {
		node := nodeMap[head]
		list = append(list, node)
		head = node.Next
	}
	return list
}

// Reverse a slice of Node pointers
func reverseList(list []*Node) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

// Merge two lists according to the rules
func mergeLists(longList, shortList []*Node) []*Node {
	var mergedList []*Node
	shortLen := len(shortList)
	longLen := len(longList)

	i, j := 0, 0
	for i < longLen {
		// Add two nodes from the long list
		if i < longLen {
			mergedList = append(mergedList, longList[i])
			i++
		}
		if i < longLen {
			mergedList = append(mergedList, longList[i])
			i++
		}
		// Add one node from the short list (in reversed order)
		if j < shortLen {
			mergedList = append(mergedList, shortList[j])
			j++
		}
	}
	return mergedList
}
