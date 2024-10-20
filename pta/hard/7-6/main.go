package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
计算岛屿最大面积
给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

岛屿的面积是岛上值为 1 的单元格的数目。

计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

示例 1：

输入：grid = [

[1, 1, 0, 0, 0];

 [1, 1, 0, 0, 0];

[0, 0, 1, 0, 0];

 [0, 0, 0, 1, 1]]

输出：4

示例 2：

输入：grid = [[0,0,0,0,0,0,0,0]]

输出：0
*/

// 深度优先搜索（DFS）计算岛屿面积
func dfs(grid [][]int, i, j int) int {
	// 如果超出边界或者当前单元格不是陆地，则返回面积为0
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}

	// 将当前单元格标记为0，避免重复访问
	grid[i][j] = 0

	// 计算当前单元格及其四周相邻的岛屿面积
	area := 1
	area += dfs(grid, i-1, j) // 上
	area += dfs(grid, i+1, j) // 下
	area += dfs(grid, i, j-1) // 左
	area += dfs(grid, i, j+1) // 右

	return area
}

// 计算最大岛屿面积
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	// 遍历整个矩阵
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				// 发现一个岛屿，启动DFS，计算面积
				area := dfs(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// 将字符串格式的二维数组解析为 [][]int
func parseGrid(input string) [][]int {
	input = strings.ReplaceAll(input, "[", "")
	input = strings.ReplaceAll(input, "]", "")
	rows := strings.Split(input, ";")
	grid := make([][]int, len(rows))

	for i, row := range rows {
		nums := strings.Split(row, ",")
		grid[i] = make([]int, len(nums))
		for j, num := range nums {
			if num == "1" {
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
		}
	}

	return grid
}

func main() {
	// 输入示例
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")

	// 解析输入
	grid := parseGrid(line)

	// 计算最大岛屿面积
	result := maxAreaOfIsland(grid)

	// 输出结果
	fmt.Println(result) // 输出: 4
}
