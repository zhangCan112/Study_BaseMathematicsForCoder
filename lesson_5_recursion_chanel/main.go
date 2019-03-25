package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

var amount = flag.Int("amount", 0, "初始金额")
var procs = flag.Int("procs", 1, "最大线程数")

var amountCategory = []int{1, 2, 5, 10}

var totalResults = make(chan [][]int, *procs)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*procs)
	start := time.Now()
	go goWork(*amount, make([]int, 0), *amount)

	for index := 0; index < count; index++ {

	}

	end := time.Now()

	fmt.Printf("您输入的金额方案有：%d种方式\n", len(results))
	fmt.Printf("计算耗时：%d秒\n", end.Unix()-start.Unix())
}

func goWork(amount int, process []int, total int) {
	if amount == 0 {
		results = append(results, process)
		fmt.Println(process)
	} else {
		for _, value := range amountCategory {
			if amount >= value {
				process := append(process, value)
				if total == amount {
					go goWork(amount-value, process, total)
				} else {
					goWork(amount-value, process, total)
				}
			}
		}
	}
}
