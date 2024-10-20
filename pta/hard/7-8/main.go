package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
解码异或后的排列

给你一个整数数组 perm ，它是前 n 个正整数（1,2,3,4,5,…,n-1,n 共n个正整数）的排列，且 n 是个奇数 。

它被加密成另一个长度为 n-1 的整数数组 encoded ，满足 encoded[i] = perm[i] XOR perm[i+1]。比方说，如果 perm=[1,3,2] ，那么 encoded=[2,1]。

给你 encoded 数组，请你返回原始数组 perm 。题目保证答案存在且唯一。

提示：

n 是奇数。

3 <= n < 10^5

encoded.length == n - 1

输入格式:
整数数组encoded，以",”分隔字符串形式作为输入

输出格式:
解码后的原始整数数组perm，以",”分隔字符串形式作为输出
*/

// 解码 perm 数组
func decode(encoded []int) []int {
	n := len(encoded) + 1
	perm := make([]int, n)

	// 计算 total = 1 XOR 2 XOR ... XOR n
	total := 0
	for i := 1; i <= n; i++ {
		total ^= i
	}

	// 计算 encoded 的部分 XOR
	oddXOR := 0
	for i := 1; i < len(encoded); i += 2 {
		oddXOR ^= encoded[i]
	}

	// 计算 perm[0]
	perm[0] = total ^ oddXOR

	// 逐个解码 perm 数组
	for i := 0; i < len(encoded); i++ {
		perm[i+1] = perm[i] ^ encoded[i]
	}

	return perm
}

// 从输入字符串解析整数数组
func parseInput(input string) []int {
	parts := strings.Split(input, ",")
	arr := make([]int, len(parts))
	for i, part := range parts {
		arr[i], _ = strconv.Atoi(part)
	}
	return arr
}

// 输出结果为逗号分隔的字符串
func formatOutput(arr []int) string {
	strs := make([]string, len(arr))
	for i, num := range arr {
		strs[i] = strconv.Itoa(num)
	}
	return strings.Join(strs, ",")
}

func main() {
	// 读取输入 encoded 数组
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")

	// 解析输入
	encoded := parseInput(line)

	// 解码 perm 数组
	perm := decode(encoded)

	// 输出解码后的 perm 数组
	fmt.Println(formatOutput(perm))
}
