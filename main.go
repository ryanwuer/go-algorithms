package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// test substring of string
	s := "hello"
	println(s[:2])
	fmt.Println(reflect.TypeOf(s[:0]))
}
