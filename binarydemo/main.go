package main

import (
	"fmt"
	"strconv"
)

/**
实现一个任意10进制数转2进制展示的方法，用位运算实现
*/

func main() {
	fmt.Println(tenToTwoBinaryString(99))
}

func tenToTwoBinaryString(num int) string {
	result := ""
	cur := num
	for cur > 0 {
		result = result + strconv.Itoa((cur)&1)
		cur = cur >> 1
	}
	return result
}
