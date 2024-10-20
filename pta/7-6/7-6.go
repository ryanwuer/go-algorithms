package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var k int
	_, err := fmt.Scanf("%d", &k)
	if err != nil {
		panic(err)
	}
	var res [][]int
	if res, err = run(k); err != nil {
		panic(err)
	}
	if len(res) == 0 {
		fmt.Println("")
		return
	}
	for i := len(res) - 1; i >= 0; i-- {
		str := intArrayToString(res[i])
		fmt.Println(str)
	}
}

func run(k int) ([][]int, error) {
	n := 2
	res := make([][]int, 0)
	for {
		added := getGapValue(n)
		if added >= k {
			break
		}
		first := (k - added) / n
		if (k-added)%n == 0 && first > 0 {
			tmp := make([]int, n)
			tmp[0] = first
			for i := 1; i < n; i++ {
				tmp[i] = tmp[i-1] + i
			}
			res = append(res, tmp)
		}
		n++
	}

	return res, nil
}

func getGapValue(n int) int {
	tmp := n
	sum := 0
	for i := 1; i < n; i++ {
		sum += (tmp - 1) * i
		tmp--
	}

	return sum
}

func intArrayToString(a []int) string {
	strArr := make([]string, len(a))
	for i, num := range a {
		strArr[i] = strconv.Itoa(num)
	}
	return strings.Join(strArr, ",")
}
