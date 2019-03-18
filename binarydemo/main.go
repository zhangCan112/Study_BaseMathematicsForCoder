package main

import (
	"fmt"
	"strconv"
)

/**
实现一个任意10进制数转2进制展示的方法，用位运算实现
*/

func main() {
	fmt.Println(tenToTwoBinaryStringForInt(-99))
}

func tenToTwoBinaryStringForUInt(num int) string {
	result := ""
	cur := num
	for cur > 0 {
		result = result + strconv.Itoa((cur)&1)
		cur = cur >> 1
	}
	return result
}

func tenToTwoBinaryStringForInt(num int) string {
	if num < 0 {
		var unum = -num
		result := ""
		cur := unum
		cur = ((cur >> 1) << 1) | ((cur & 1) ^ 1)
		for cur > 0 {
			result = result + strconv.Itoa((((cur) & 1) ^ 1))
			cur = cur >> 1
		}
		return result
	}
	return tenToTwoBinaryStringForUInt(num)
}
