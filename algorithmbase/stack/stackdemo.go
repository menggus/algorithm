package main

import "fmt"

func main() {
	// 创建一个栈
	stack := make([]int, 0)

	// 入栈
	stack = append(stack, 7)

	// 出栈
	_ = stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	// 栈空检测
	fmt.Println(len(stack) == 0)
}
