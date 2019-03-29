package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

var amount = flag.Int("amount", 0, "初始金额")
var procs = flag.Int("procs", 8, "最大线程数")

var amountCategory = []int{1, 2, 5, 10}

var resultChanel = make(chan []int, *procs)
var refChangeChanel = make(chan int, *procs)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*procs)
	start := time.Now()
	go goWork(*amount, make([]int, 0), *amount)

	results := make([][]int, 0)
	refCount := 1
	for refCount > 0 {
		select {
		case process := <-resultChanel:
			results = append(results, process)
			fmt.Println(process)
		case change := <-refChangeChanel:
			{
				refCount += change
				if refCount == 1 {
					refChangeChanel <- -1
					close(refChangeChanel)
					close(resultChanel)
				}
			}
		default:
			// fmt.Println("wait result!")
		}
	}

	end := time.Now()

	fmt.Printf("您输入的金额方案有：%d种方式\n", len(results))
	fmt.Printf("计算耗时：%d秒\n", end.Unix()-start.Unix())
}

func goWork(amount int, process []int, total int) {

	refChangeChanel <- 1
	if amount == 0 {
		resultChanel <- process
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
	time.Sleep(1e7)
	refChangeChanel <- -1
}
