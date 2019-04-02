/*
给定一个字段，实现其前缀树的生成与查询参数
*/
package main

import "fmt"

var originMap = map[string]string{
	"a":         "一（个）;每一（个）;任一（个）",
	"and":       "而且;和，与;于是，然后;因此",
	"as":        "由于;像，像…一样;同时，当…时;尽管",
	"above":     "表示程度）超过;（表示等级）在…之上;（表示位置）在…正上方;（表示比较）优于",
	"ability":   "能力，资格;能耐，才能",
	"about":     "关于;大约;在…周围",
	"accept":    "接受;承认;承担;承兑",
	"account":   "账，账目;存款;记述，报告;理由",
	"actually":  "实际上;确实;竟;事实上",
	"ad":        "广告;网球中的优势分advantage. ad in",
	"address":   "地址;称呼;演说;通信处",
	"advice":    "建议;劝告，忠告;（商业）通知;（政治，外交上的）报导，报告",
	"advance":   "预付;提出;（使）前进;将…提前",
	"bear":      "熊;（在证券市场等）卖空的人;蛮横的人",
	"beautiful": "beautiful",
	"history":   "历史",
	"available": "可获得的",
	"break":     "突破;（嗓音）突变;破晓;（价格）突然下跌",
	"correct":   "正确的;合适的;符合公认准则的;得体的",
	"dress":     "连衣裙的;须穿礼服的;适合于正式场合的;办公时（或半正式场合）穿戴的",
}

type TreeNode struct {
	label   string
	sons    map[string]TreeNode
	explain string
}

func main() {
	fmt.Println(originMap["history"])
	fmt.Println(originMap["histor1y"])
}
