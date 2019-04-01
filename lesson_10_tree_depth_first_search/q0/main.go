/*
给定一个字段，实现其前缀树的生成与查询参数
*/
package main

import "fmt"

var originMap = map[string]string{
	"history":   "历史",
	"available": "可获得的",
	"break":     "突破;（嗓音）突变;破晓;（价格）突然下跌",
	"correct":   "正确的;合适的;符合公认准则的;得体的",
	"dress":     "连衣裙的;须穿礼服的;适合于正式场合的;办公时（或半正式场合）穿戴的",
}

func main() {
	fmt.Println(originMap["history"])
	fmt.Println(originMap["histor1y"])
}
