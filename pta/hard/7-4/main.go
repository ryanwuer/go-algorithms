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
最长递增子序列

给你一个整数数组nums，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

输入格式:
1 <= nums.length <= 2000
-10000 <= nums[i] <= 10000

输出格式:
最长严格递增子序列的长度
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	str := strings.Split(line, " ")
	var array = make([]int, 0)
	for i := 0; i < len(str); i++ {
		v, _ := strconv.Atoi(str[i])
		array = append(array, v)
	}
	var dp = make([]int, len(str))
	dp[0] = 1
	for i := 1; i < len(array); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if array[j] < array[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	sort.Slice(dp, func(i, j int) bool {
		return dp[i] > dp[j]
	})

	fmt.Println(dp[0])
}
