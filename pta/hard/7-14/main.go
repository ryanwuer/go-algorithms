package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
按位与为零的三元组

给你一个整数数组 nums ，返回其中 按位与三元组 的数目。

按位与三元组 是由下标 (i, j, k) 组成的三元组，并满足下述全部条件：

0 <= i < nums.length

0 <= j < nums.length

0 <= k < nums.length

nums[i] & nums[j] & nums[k] == 0 ，其中 & 表示按位与运算符。

提示：

1 <= nums.length <= 1000

0 <= nums[i] < 2^16

说明：注意算法复杂度，暴力枚举解法，会超出时间限制

输入格式:
整数数组nums，以",”分隔字符串形式作为输入

输出格式:
一个数字，字符串形式
*/

// 计算按位与三元组的数目
func countTriplets(nums []int) int {
	count := make(map[int]int) // 用来记录每个 nums[i] & nums[j] 结果的出现次数
	n := len(nums)

	// 计算所有 nums[i] & nums[j] 并记录到 count 中
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			andResult := nums[i] & nums[j]
			count[andResult]++
		}
	}

	result := 0

	// 枚举 nums[k]，并在 count 中查找有多少 nums[i] & nums[j] & nums[k] == 0
	for _, num := range nums {
		for key, val := range count {
			if num&key == 0 {
				result += val
			}
		}
	}

	return result
}

func main() {
	// 从终端读取输入
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// 去掉换行符
	input = strings.TrimRight(input, "\n")

	// 将输入的字符串转换为整数数组
	numStrs := strings.Split(input, ",")
	nums := make([]int, len(numStrs))
	for i, str := range numStrs {
		num, _ := strconv.Atoi(str)
		nums[i] = num
	}

	// 计算按位与三元组的数目
	result := countTriplets(nums)

	// 输出结果
	fmt.Println(result)
}
