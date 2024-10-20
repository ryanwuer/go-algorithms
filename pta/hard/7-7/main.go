package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
不相交的线

在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。

现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：

nums1[i] == nums2[j]

且绘制的直线不与任何其他连线（非水平线）相交。

请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。

以这种方法绘制线条，并返回可以绘制的最大连线数。

1 <= nums1.length, nums2.length <= 500

1 <= nums1[i], nums2[j] <= 2000

输入格式:
每组输入为两行，表示nums1和nums2两个数组。每行有n+1个数字，数字间用空格分开，第一个数字表示数组个数n，后面跟n个数字；如2 2 3，表示数组有2个元素，元素值为2和3

输出格式:
输出最多能绘制不想交线的条数。
*/

// 计算最大连线数
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// dp数组初始化
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	// 动态规划求解
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}

// 求两个数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解析输入
func parseLine(line string) []int {
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0])
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i+1])
	}
	return nums
}

func main() {
	// 输入示例
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	// 解析输入
	nums1 := parseLine(line)
	line, _ = reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	// 解析输入
	nums2 := parseLine(line)

	// 计算最大连线数
	result := maxUncrossedLines(nums1, nums2)

	// 输出结果
	fmt.Println(result) // 输出: 2
}
