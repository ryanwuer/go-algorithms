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
二叉树的最大路径和

二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

输入格式:
树上的节点数满足 0 <= n <= 1000, 每个节点的值满足 -1000 <= val <= 1000

（注：null节点的左右叶子不会再打印null）

输出格式:
输出最大路径的和
*/

// 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量存储最大路径和
var maxSum int

// 递归函数：计算节点的最大贡献值
func maxGain(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// 递归计算左右子树的最大贡献值
	leftGain := max(maxGain(node.Left), 0)
	rightGain := max(maxGain(node.Right), 0)

	// 计算当前节点的最大路径和
	currentMaxPath := node.Val + leftGain + rightGain

	// 更新全局最大路径和
	maxSum = max(maxSum, currentMaxPath)

	// 返回当前节点的最大贡献值
	return node.Val + max(leftGain, rightGain)
}

// 主函数，计算最大路径和
func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt64 // 初始化最大路径和为最小值
	maxGain(root)          // 递归计算最大路径和
	return maxSum
}

// 工具函数：计算较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 工具函数：根据输入构建二叉树
func buildTree(nodes []string) *TreeNode {
	if len(nodes) == 0 || nodes[0] == "null" {
		return nil
	}

	val, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}

	i := 1
	for i < len(nodes) {
		curr := queue[0]
		queue = queue[1:]

		// 左边节点
		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])
			curr.Left = &TreeNode{Val: val}
			queue = append(queue, curr.Left)
		}
		i++

		// 右边节点
		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])
			curr.Right = &TreeNode{Val: val}
			queue = append(queue, curr.Right)
		}
		i++
	}

	return root
}

func main() {
	// 输入数据
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	nodes := strings.Split(line, ",")
	root := buildTree(nodes)

	// 计算并输出最大路径和
	fmt.Println(maxPathSum(root)) // 输出: 42
}
