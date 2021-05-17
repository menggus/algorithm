package main

import "fmt"

// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1

func main() {
	fmt.Println(SearchFirstStrPosition("你好呀abc飒飒休闲", "飒休"))
	fmt.Println(SearchFirstStrPosition2("你好呀abc飒飒休闲", "飒休"))
}

func SearchFirstStrPosition(haystack, needle string) int {
	//  abcdefg
	//  cde
	if len(needle) == 0 {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// for 循环对比每一个字符来判断是否相等
		for n := 0; n < len(needle); n++ {
			if haystack[i+n] != needle[n] {
				break
			}
			if n == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

func SearchFirstStrPosition2(haystack, needle string) int {
	//  abcdefg
	//  cde
	if len(needle) == 0 {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// 这里判断直接采用的是 字符串 判等。经benchmark测试效率要低于 for 循环
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
