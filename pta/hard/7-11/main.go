package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	// cnt is a map to keep track of the count of elements in the original array and the sorted array.
	cnt := map[int]int{}

	// b is a copy of the original array arr, which is then sorted.
	b := append([]int{}, arr...)
	sort.Ints(b)

	// Iterate through the original array.
	for i, x := range arr {
		// Increment the count for the current element in the original array.
		cnt[x]++
		// If the count becomes zero, remove the element from the map.
		if cnt[x] == 0 {
			delete(cnt, x)
		}

		// Get the corresponding element from the sorted array.
		y := b[i]
		// Decrement the count for the current element in the sorted array.
		cnt[y]--
		// If the count becomes zero, remove the element from the map.
		if cnt[y] == 0 {
			delete(cnt, y)
		}

		// If the map is empty, it means we can form a chunk.
		if len(cnt) == 0 {
			ans++
		}
	}
	return
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
