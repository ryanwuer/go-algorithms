package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 判断一个数是否是回文数
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	rev := reverseString(s)
	return s == rev
}

// 反转字符串
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 生成在范围 maxNum 内的所有可能的回文数
func generatePalindromes(maxNum int) []int {
	var result []int

	for i := 1; i < 100000; i++ {
		s := strconv.Itoa(i)
		r := reverseString(s)
		abba, _ := strconv.Atoi(s + r)
		if abba <= maxNum {
			result = append(result, abba)
		}
		r = reverseString(s[:len(s)-1])
		aba, _ := strconv.Atoi(s + r)
		if aba <= maxNum {
			result = append(result, aba)
		}
	}

	return result
}

// 找到在给定范围内的所有超级回文数
func superPalindromesInRange(L, R int) []int {
	var result []int

	maxNum := int(math.Sqrt(float64(R))) + 1
	palindromes := generatePalindromes(maxNum)
	for i := 0; i < len(palindromes); i++ {
		tmp := palindromes[i] * palindromes[i]
		flag := isPalindrome(tmp)
		if tmp >= L && tmp <= R && flag {
			result = append(result, tmp)
		}
	}

	sort.Ints(result)

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n")
	nums := strings.Split(line, ",")
	L, _ := strconv.Atoi(nums[0])
	R, _ := strconv.Atoi(nums[1])
	result := superPalindromesInRange(L, R)
	// print [4,9,121,484]
	resultStr := make([]string, len(result))
	for i, r := range result {
		resultStr[i] = strconv.Itoa(r)
	}
	fmt.Println("[" + strings.Join(resultStr, ", ") + "]")
}
