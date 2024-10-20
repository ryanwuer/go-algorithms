package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// 查看栈顶元素但不弹出
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

// 查看栈是否为空
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

type MyQueue struct {
	pushStack Stack
	popStack  Stack
}

// Constructor 初始化一个新的 MyQueue 对象
func Constructor() MyQueue {
	return MyQueue{
		pushStack: Stack{},
		popStack:  Stack{},
	}
}

// Push 将元素 x 入队
func (q *MyQueue) Push(x int) {
	q.pushStack.Push(x)
}

// Pop 从队列中移除并返回第一个元素
func (q *MyQueue) Pop() int {
	// 如果 popStack 为空，将 pushStack 中的元素倒入 popStack
	if q.popStack.Empty() {
		for !q.pushStack.Empty() {
			top := q.pushStack.Pop()
			q.popStack.Push(top)
		}
	}
	return q.popStack.Pop()
}

// Peek 返回队列中的第一个元素
func (q *MyQueue) Peek() int {
	// 如果 popStack 为空，将 pushStack 中的元素倒入 popStack
	if q.popStack.Empty() {
		for !q.pushStack.Empty() {
			top := q.pushStack.Pop()
			q.popStack.Push(top)
		}
	}
	return q.popStack.Peek()
}

// Empty 检查队列是否为空
func (q *MyQueue) Empty() bool {
	return q.pushStack.Empty() && q.popStack.Empty()
}

func main() {
	q := Constructor()
	// 输入两行内容，第一行表示各种操作，逗号分隔，第二行表示操作的元素，逗号分隔，如果操作没有参数，则不输入
	// 第一行举例：push,push,peek,pop,empty
	// 第二行举例：1,2,,,
	// 输出举例：nil,nil,1,1,false
	var ops, params string
	_, err := fmt.Scanln(&ops)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanln(&params)
	if err != nil {
		panic(err)
	}
	oItems := strings.Split(ops, ",")
	pItems := strings.Split(params, ",")
	res := make([]interface{}, len(oItems))
	for i, o := range oItems {
		p := pItems[i]
		switch o {
		case "push":
			item, _ := strconv.Atoi(p)
			q.Push(item)
			res[i] = nil
		case "pop":
			res[i] = q.Pop()
		case "peek":
			res[i] = q.Peek()
		case "empty":
			res[i] = q.Empty()
		default:
			res[i] = nil
		}
	}
	// print res with items joined by ,
	for i, r := range res {
		if r == nil {
			fmt.Print("null")
		} else {
			fmt.Print(r)
		}
		if i < len(res)-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("\n")
}
