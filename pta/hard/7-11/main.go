package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
最多能完成排序的块

给你一个整数数组 arr 。将 arr 分割成若干 块 ，并将这些块分别进行排序。之后再连接起来，使得连接的结果和按升序排序后的原数组相同。

返回能将数组分成的最多块数？

输入格式:
1. 输入整数数列，元素之间以空格分开

2. 其中数组长度为n，1<=n<=1000,

3. 数组元素 1 <= arr[i], k <= 100，数组元素可有重复整数

输出格式:
数组能分成的最多块数
*/

// 计算数组能分成的最多块数
// maxChunksToSorted calculates the maximum number of chunks that the array can be split into
// such that sorting each chunk individually results in the entire array being sorted.
func maxChunksToSorted(arr []int) (ans int) {
	var st []int // 用来存储块
	for _, x := range arr {
		// 如果当前元素比栈顶大，直接入栈
		if len(st) == 0 || x >= st[len(st)-1] {
			st = append(st, x)
		} else {
			// 当前元素比栈顶小，开始合并块
			mx := st[len(st)-1] // 记录当前栈顶最大值
			st = st[:len(st)-1] // 弹出栈顶

			// 把比当前元素 x 大的栈顶元素都弹出
			for len(st) > 0 && st[len(st)-1] > x {
				st = st[:len(st)-1]
			}
			// 放入当前块的最大值 mx，保持块的正确顺序
			st = append(st, mx)
		}
	}
	return len(st) // 栈中元素的数量就是块的数量
}

func main() {
	// 使用 bufio 来读取标准输入
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// 将输入的字符串分割成数组
	parts := strings.Fields(input)
	arr := make([]int, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		arr[i] = num
	}

	// 计算并输出最多能分成的块数
	result := maxChunksToSorted(arr)
	fmt.Println(result)
}
