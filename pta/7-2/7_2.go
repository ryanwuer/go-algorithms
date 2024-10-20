package main

import (
	"fmt"
)

// DFS函数，用于遍历相邻的山峰，并标记它们为已访问
func dfs(grid [][]int, visited [][]bool, x, y, M, N int) {
	// 定义四个方向：上、下、左、右
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 标记当前山峰为已访问
	visited[x][y] = true

	// 遍历四个方向
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]

		// 检查边界，并且确认相邻点是山峰且未访问
		if newX >= 0 && newX < M && newY >= 0 && newY < N && grid[newX][newY] == 1 && !visited[newX][newY] {
			dfs(grid, visited, newX, newY, M, N)
		}
	}
}

func countMountains(grid [][]int, M, N int) int {
	// 初始化一个已访问的标记矩阵
	visited := make([][]bool, M)
	for i := range visited {
		visited[i] = make([]bool, N)
	}

	mountainCount := 0

	// 遍历整个地图
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			// 如果当前是一个山峰且未访问过
			if grid[i][j] == 1 && !visited[i][j] {
				// 进行一次DFS，标记整个大山
				dfs(grid, visited, i, j, M, N)
				// 每次DFS完毕，说明找到一个完整的大山
				mountainCount++
			}
		}
	}

	return mountainCount
}

func main() {
	// 输入地图的行数和列数
	var M, N int
	fmt.Scan(&M, &N)

	// 输入地图
	grid := make([][]int, M)
	for i := range grid {
		grid[i] = make([]int, N)
		for j := range grid[i] {
			fmt.Scan(&grid[i][j])
		}
	}

	// 计算大山数目
	result := countMountains(grid, M, N)

	// 输出大山的数目
	fmt.Println(result)
}
