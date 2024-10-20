package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。


输入样例1:
在这里给出一组输入。例如：

abcabcbb
输出样例1:
在这里给出相应的输出。例如：

3
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	var array []rune
	m := map[rune]int{}
	maxLen := 0
	for i, v := range line {
		if j, ok := m[v]; ok {
			// 存在的话，更新当前 maxLen，并弹出前几个元素
			maxLen = max(maxLen, i-j)
		} else {
			array = append(array, v)
		}
		// 无论如何，都更新字符 a 的最新下标位置
		m[v] = i
	}

	fmt.Println(len(array))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
