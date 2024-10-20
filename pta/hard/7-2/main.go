package main

import (
	"fmt"
	"math"
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

// 并查集结构体
type UnionFind struct {
	parent []int
	size   []int
}

// 初始化并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent, size}
}

// 查找操作，找到 x 的根
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// 合并操作，将 x 和 y 合并
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		// 合并时根据 size 优化
		if uf.size[rootX] > uf.size[rootY] {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
		} else {
			uf.parent[rootX] = rootY
			uf.size[rootY] += uf.size[rootX]
		}
	}
}

// 找到最大连通分量的大小
func (uf *UnionFind) MaxSize() int {
	maxSize := 0
	for i := 0; i < len(uf.size); i++ {
		if uf.size[i] > maxSize {
			maxSize = uf.size[i]
		}
	}
	return maxSize
}

// 质因数分解和并查集连接
func largestComponentSize(nums []int) int {
	// 找到数组中的最大值
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	// 初始化并查集
	uf := NewUnionFind(maxNum + 1)

	// 对每个数进行质因数筛选
	for _, num := range nums {
		for factor := 2; factor <= int(math.Sqrt(float64(num))); factor++ {
			if num%factor == 0 {
				uf.Union(num, factor)
				uf.Union(num, num/factor)
			}
		}
	}

	// 统计每个连通分量的大小
	count := make(map[int]int)
	for _, num := range nums {
		root := uf.Find(num)
		count[root]++
	}

	// 找到最大的连通分量
	maxSize := 0
	for _, size := range count {
		if size > maxSize {
			maxSize = size
		}
	}

	return maxSize
}

// 主函数
func main() {
	// 从输入读取
	var nums []int
	for {
		var num int
		n, err := fmt.Scan(&num)
		if n == 0 || err != nil {
			break
		}
		nums = append(nums, num)
	}

	// 计算并输出最大连通分量的大小
	fmt.Println(largestComponentSize(nums))
}
