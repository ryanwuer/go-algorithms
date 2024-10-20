package main

import (
	"fmt"
)

/*
拼接最大数

给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。

现在从这两个数组中选出 k (0 <=k <= m + n) 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。

输入格式:
输入三个行内容：

第一行是数组nums1，元素内容用逗号分隔；数组最大长度为1000。

第二行是数组nums2，元素内容用逗号分隔；数组最大长度为1000。

第三行是长度k；

输出格式:
返回一个表示该最大数的长度为 k 的数组，数组元素用逗号隔开。
*/

// 获取数组中挑选长度为 t 的最大子序列
func maxSubsequence(nums []int, t int) []int {
	stack := make([]int, 0, t)
	drop := len(nums) - t

	for _, num := range nums {
		// 贪心算法，丢弃较小的元素
		for len(stack) > 0 && drop > 0 && stack[len(stack)-1] < num {
			stack = stack[:len(stack)-1]
			drop--
		}
		stack = append(stack, num)
	}

	return stack[:t] // 保留前 t 个元素
}

// 合并两个数组，构建最大数
func merge(nums1, nums2 []int) []int {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if compare(nums1, i, nums2, j) {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	return merged
}

// 比较两个数组的字典序大小
func compare(nums1 []int, i int, nums2 []int, j int) bool {
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] != nums2[j] {
			return nums1[i] > nums2[j]
		}
		i++
		j++
	}
	return len(nums1)-i > len(nums2)-j
}

// 主函数，计算最终最大数
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	maxResult := make([]int, k)

	// 尝试在 nums1 中取 i 个元素，nums2 中取 k-i 个元素
	for i := 0; i <= k && i <= len(nums1); i++ {
		if k-i <= len(nums2) {
			subseq1 := maxSubsequence(nums1, i)
			subseq2 := maxSubsequence(nums2, k-i)
			merged := merge(subseq1, subseq2)
			if compare(merged, 0, maxResult, 0) {
				copy(maxResult, merged)
			}
		}
	}

	return maxResult
}

// 主函数，处理输入和输出
func main() {
	// 输入读取
	var nums1Str, nums2Str string
	var k int
	fmt.Scan(&nums1Str)
	fmt.Scan(&nums2Str)
	fmt.Scan(&k)

	// 解析输入数组
	var nums1, nums2 []int
	for _, num := range split(nums1Str) {
		nums1 = append(nums1, num)
	}
	for _, num := range split(nums2Str) {
		nums2 = append(nums2, num)
	}

	// 计算结果并输出
	result := maxNumber(nums1, nums2, k)
	for i, num := range result {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(num)
	}
	fmt.Println()
}

// 工具函数：将字符串解析为整数数组
func split(s string) []int {
	var res []int
	var num int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == ',' {
			res = append(res, num)
			num = 0
		}
	}
	res = append(res, num)
	return res
}
