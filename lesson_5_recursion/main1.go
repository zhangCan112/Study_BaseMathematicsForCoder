package main

import (
	"flag"
	"fmt"
	"time"
)

var amount = flag.Int("amount", 0, "初始金额")

var amountCategory = []int{1, 2, 5, 10}

var results = make([][]int, 0)

func main() {
	flag.Parse()
	start := time.Now()
	goWork(*amount, make([]int, 0))
	end := time.Now()
	fmt.Printf("您输入的金额方案有：%d种方式\n", len(results))
	fmt.Printf("计算耗时：%d秒\n", end.Unix()-start.Unix())
}

func goWork(amount int, process []int) {
	time.Sleep(1e7)
	if amount == 0 {
		fmt.Println(process)
		results = append(results, process)
	} else {
		for _, value := range amountCategory {
			if amount >= value {
				process := append(process, value)
				goWork(amount-value, process)
			}
		}
	}
}
