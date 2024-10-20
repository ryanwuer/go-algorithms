package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	res := run(n)
	fmt.Println(res)
}

func run(n int) int {
	// 5, 2, 1
	sum := 0

	fiveCount := n / 5
	sum += fiveCount
	n = n % 5
	if n == 0 {
		return sum
	}
	twoCount := n / 2
	sum += twoCount
	n = n % 2
	if n == 0 {
		return sum
	}
	return sum + n
}
