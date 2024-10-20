package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	if len(str) > 2000 {
		fmt.Println(0)
		return
	}

	s := strings.ToLower(str)

	// 移除非字母数字字符
	filtered := ""
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			filtered += string(char)
		}
	}

	if len(filtered) == 0 {
		fmt.Println(0)
		return
	}

	// 检查是否是回文串
	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			fmt.Println(0)
			return
		}
		left++
		right--
	}
	fmt.Println(1)
}
