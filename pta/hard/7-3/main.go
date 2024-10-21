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

func maxNumber(nums1, nums2 []int, k int) []int {
	// 从一个数组中取 k 个数，保持相对顺序，并且结果最大
	maxSubArray := func(nums []int, k int) []int {
		stack := []int{}
		drop := len(nums) - k
		for _, num := range nums {
			for drop > 0 && len(stack) > 0 && num > stack[len(stack)-1] {
				// 当前数大于栈顶元素，则弹出，并且 drop 减去一
				stack = stack[:len(stack)-1]
				drop--
			}
			// num 压入栈
			stack = append(stack, num)
		}
		// 返回前 k 个元素
		return stack[:k]
	}

	// 合并两个数组成最大值
	merge := func(nums1, nums2 []int) []int {
		res := []int{}
		for len(nums1) > 0 || len(nums2) > 0 {
			if greater(nums1, nums2) {
				res = append(res, nums1[0])
				nums1 = nums1[1:]
			} else {
				res = append(res, nums2[0])
				nums2 = nums2[1:]
			}
		}
		return res
	}

	// 主函数逻辑，枚举可能的分配方式
	res := []int{}
	for i := 0; i <= k && i <= len(nums1); i++ {
		if k-i <= len(nums2) {
			sub1 := maxSubArray(nums1, i)
			sub2 := maxSubArray(nums2, k-i)
			res = max(res, merge(sub1, sub2))
		}
	}
	return res
}

// 比较两个数组字典序大小
func greater(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] > nums2[i] {
			return true
		}
		if nums1[i] < nums2[i] {
			return false
		}
	}
	return len(nums1) > len(nums2)
}

// 比较两个数组哪个更大
func max(nums1, nums2 []int) []int {
	if greater(nums1, nums2) {
		return nums1
	}
	return nums2
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
