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

func tenToTwoBinaryStringForInt(num int) string {
	result := ""
	arr := make([]int, 0, 1)
	cur := num
	if num < 0 {
		cur = -num
	}

	for cur > 0 {
		arr = append(arr, (cur)&1)
		cur = cur >> 1
	}

	if num < 0 {
		arr = reverse(arr)
		arr = addOne(arr)
	}

	for _, bit := range arr {
		result += strconv.Itoa(bit)
	}
	return result
}

func reverse(arr []int) []int {
	var newArr = make([]int, 0, 0)
	for _, bit := range arr {
		newArr = append(newArr, ^bit&1)
	}
	return newArr
}

func addOne(arr []int) []int {
	var need = true
	for index := len(arr) - 1; index > -1 && need; index-- {
		need = (arr[index]&1 == 1)
		arr[index] = arr[index] | 1
	}
	return arr
}
