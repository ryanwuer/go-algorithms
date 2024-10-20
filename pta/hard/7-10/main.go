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

/*
超级回文数

如果一个正整数自身是回文数，而且它也是一个回文数的平方，那么我们称这个数为超级回文数。

现在，给定两个正整数 L 和 R ，请按照从小到大的顺序打印包含在范围 [L, R] 中的所有超级回文数。

注：R包含的数字不超过20位

回文数定义：将该数各个位置的数字反转排列，得到的数和原数一样，例如676，2332，10201。

输入格式:
L,R。例如4,1000

输出格式:
[L, R]范围内的超级回文数，例如[4, 9, 121, 484]
*/

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
	var palindromes []int
	// 生成形式为 'abba' 和 'aba' 的回文数
	for i := 1; i < 100000; i++ {
		s := strconv.Itoa(i)
		// 偶数长度回文数
		rev := reverseString(s)
		abbapalindrome, _ := strconv.Atoi(s + rev)
		if abbapalindrome <= maxNum {
			palindromes = append(palindromes, abbapalindrome)
		}
		// 奇数长度回文数
		rev = reverseString(s[:len(s)-1])
		abapalindrome, _ := strconv.Atoi(s + rev)
		if abapalindrome <= maxNum {
			palindromes = append(palindromes, abapalindrome)
		}
	}
	return palindromes
}

// 找到在给定范围内的所有超级回文数
func superPalindromesInRange(L, R int) []int {
	var result []int
	maxSqrt := int(math.Sqrt(float64(R))) + 1
	palindromes := generatePalindromes(maxSqrt)

	for _, p := range palindromes {
		pSquared := p * p
		if pSquared >= L && pSquared <= R && isPalindrome(pSquared) {
			result = append(result, pSquared)
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
