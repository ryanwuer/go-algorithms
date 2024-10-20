package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	ks, _ := reader.ReadString('\n')
	ks = strings.TrimSpace(ks)
	k, _ := strconv.Atoi(ks)

	strNums := strings.Split(input, " ")
	numbers := make([]int, len(strNums))

	// 遍历字符串数组并将每个元素转换为整数
	for i, str := range strNums {
		num, err := strconv.Atoi(str) // 转换为整数
		if err != nil {
			return
		}
		numbers[i] = num // 存入数组
	}

	fmt.Println(k + 1 + (numbers[k]-1)*len(numbers))
}
