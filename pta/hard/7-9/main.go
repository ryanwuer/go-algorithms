package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
最长超赞子字符串

给你一个字符串 s 。请返回 s 中最长的 超赞子字符串 的长度。
「超赞子字符串」需满足满足下述两个条件：
该字符串是 s 的一个非空子字符串
进行任意次数的字符交换后，该字符串可以变成一个回文字符串
1 <= s.length <= 10^5
s 仅由数字组成

输入格式:
输入一行只包含数字的字符串s

输出格式:
输出s中最长的 超赞子字符串 的长度
*/

func longestAwesome(s string) int {
	mask, maxLen := 0, 0
	seen := map[int]int{0: -1}

	for i := 0; i < len(s); i++ {
		mask ^= 1 << (s[i] - '0')
		if pos, ok := seen[mask]; ok {
			maxLen = max(maxLen, i-pos)
		} else {
			seen[mask] = i
		}
		for j := 0; j < 10; j++ {
			if pos, ok := seen[mask^(1<<j)]; ok {
				maxLen = max(maxLen, i-pos)
			}
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 读取输入 encoded 数组
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	fmt.Println(longestAwesome(line))
}
