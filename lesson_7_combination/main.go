/**
题目： 50人中抽取一等奖1名，二等奖3名，三等奖10名，有多少种组合方式
*/

package main

import (
	"flag"
	"fmt"
)

var total = flag.Int("total", 50, "参与抽奖的人数，默认100人")

func main() {
	flag.Parse()
	fmt.Println(drawLucky([]int{1, 3, 10}, *total))
}

func drawLucky(luckers []int, total int) int {
	cur := total
	category := 1
	for _, count := range luckers {
		category *= combination(count, cur)
		cur -= count
	}
	return category
}

func combination(num int, total int) int {
	var a = 1
	var b = 1
	for index := 0; index < num; index++ {
		a *= (total - index)
		b *= (num - index)
	}
	return a / b
}
