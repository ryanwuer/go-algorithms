package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
最长有效括号

给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

输入格式:
包含 '(' 和 ')' 的字符串

输出格式:
有效括号子串的长度
*/

// 返回最长的有效括号子串的长度
func longestValidParentheses(s string) int {
	st := []int{-1} // 初始化栈，加入 -1 作为基准
	maxLen := 0     // 最大长度初始化为0

	for i, ch := range s {
		if ch == '(' {
			// 遇到左括号，把索引压入栈
			st = append(st, i)
		} else {
			// 遇到右括号，弹出栈顶
			st = st[:len(st)-1]
			if len(st) > 0 {
				// 如果栈不为空，计算当前长度
				maxLen = max(maxLen, i-st[len(st)-1])
			} else {
				// 栈为空，更新基准索引
				st = append(st, i)
			}
		}
	}

	return maxLen
}

// 返回两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 从终端读取输入
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")

	// 计算最长有效括号子串的长度
	result := longestValidParentheses(input)

	// 输出结果
	fmt.Println(result)
}
