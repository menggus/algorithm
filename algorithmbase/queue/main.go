package main

import "fmt"

func main() {
	// 创建一个队列
	queue := make([]int, 0)

	// 入队
	queue = append(queue, 8)

	// 出队
	_ = queue[0]
	queue = queue[1:]

	// 空队列检测
	fmt.Println(len(queue) == 0)
}
