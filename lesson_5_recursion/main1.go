package main

import (
	"flag"
	"fmt"
)

var amount = flag.Int("amount", 0, "初始金额")

var amountCategory = []int{1, 2, 5, 10}

func main() {
	flag.Parse()
	fmt.Printf("您输入的金额：%d\n", *amount)
}

func goWork(amount int, result int) int {
	count := result
	if amount == 0 {
		return count
	} else {
		for _, value := range amountCategory {
			if count >= value {

			}
		}
	}
}
