package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
按公因数计算最大组件大小

给定一个由不同正整数组成的非空数组 nums，考虑下面的构图：

1. 有 nums.length 个节点，按照从 nums[0]到 nums[nums.length-1]标记；

2. 只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j] 之间才有一条边。

返回构图中最大连通组件的大小。

输入格式:
输入为数组元素，空格分隔

输出格式:
输出最大连通组件的大小
*/

// 查找节点的根节点
func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

// 合并两个节点
func union(parent []int, x, y int) {
	rootX, rootY := find(parent, x), find(parent, y)
	if rootX != rootY {
		parent[rootY] = rootX
	}
}

func largestComponentSize(nums []int) int {
	maxNum := 0
	// 更新最大数
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	// 初始化 parent 数组，注意 size 等于 maxNum+1
	parent := make([]int, maxNum+1)
	for i := range parent {
		// 初始化 parent 数组，每个元素的 parent 都是自己
		parent[i] = i
	}

	for _, num := range nums {
		for factor := 2; factor <= int(math.Sqrt(float64(num))); factor++ {
			if num%factor == 0 {
				union(parent, num, factor)
				union(parent, num, num/factor)
			}
		}
	}

	// 计算每个连通分量的大小
	count := make(map[int]int)
	res := 0
	for _, num := range nums {
		root := find(parent, num)
		count[root]++
		if count[root] > res {
			res = count[root]
		}
	}
	return res
}

// 主函数
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// 去掉换行符
	input = strings.TrimRight(input, "\n")

	// 将输入的字符串转换为整数数组
	numStrs := strings.Split(input, " ")
	nums := make([]int, len(numStrs))
	for i, str := range numStrs {
		num, _ := strconv.Atoi(str)
		nums[i] = num
	}

	// 计算并输出最大连通分量的大小
	fmt.Println(largestComponentSize(nums))
}
