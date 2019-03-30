/**
动态规划练习题。 有2，3，7三种钱币，求有最少的钱币凑满100元
分别使用 迭代，递归 和动态规划的思路解决
*/

package main

import "fmt"

//金额类别
var amountCategory = [3]int{2, 3, 7}

func main() {
	//递归法调用
	fmt.Println(recursion(60))
}

/*
迭代法
*/
func iteration(amount int) []int {
	return nil
}

/*
递归法
*/
func recursion(amount int) []int {

	if amount == 0 {
		return make([]int, 0)
	}

	if amount < 0 {
		return nil
	}

	first := recursion(amount - 2)
	if first != nil {
		first = append(first, 2)
	}

	second := recursion(amount - 3)
	if second != nil {
		second = append(second, 3)
	}

	third := recursion(amount - 7)
	if third != nil {
		third = append(third, 7)
	}

	min := minArr(first, second)
	min = minArr(min, third)
	return min
}

func minArr(a []int, b []int) []int {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if len(a) < len(b) {
		return a
	}
	return b
}

/*
动态规划法
*/
func dynamicProgramming() {

}
