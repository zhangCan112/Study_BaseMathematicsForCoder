package main

import (
	"flag"
	"fmt"
)

var word = flag.String("word", "", "需要查询的单词")

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
	"beautiful": "美丽的",
	"history":   "历史",
	"available": "可获得的",
	"break":     "突破;（嗓音）突变;破晓;（价格）突然下跌",
	"correct":   "正确的;合适的;符合公认准则的;得体的",
	"dress":     "连衣裙的;须穿礼服的;适合于正式场合的;办公时（或半正式场合）穿戴的",
}

func main() {
	flag.Parse()
	fmt.Printf("要查询到的单词是%s\n", *word)
	//1.构建前缀树
	treeNode := build(originMap)

	//2. 查询执行word
	fmt.Printf("查询到的结果是\"%s\"\n", treeNode.query(*word))

	//3. 遍历字典，深度优先
	treeNode.traverseAllLeafNodes()
}

type stack struct {
	contaner []*node
}

func (s *stack) push(n *node) {
	s.contaner = append(s.contaner, n)
}

func (s *stack) pop() *node {
	if len(s.contaner) == 0 {
		return nil
	}
	result := s.contaner[len(s.contaner)-1]
	s.contaner = s.contaner[0 : len(s.contaner)-1]
	return result
}

type node struct {
	word    rune
	explain string
	sons    []*node
}

func build(dic map[string]string) *node {
	root := &node{}

	for word, explain := range dic {
		root.insertWord(word, explain)
	}
	return root
}

func (n *node) insertWord(word, explain string) {
	wordArr := []rune(word)
	n.findOrCreate(wordArr, explain)
}

func (n *node) findOrCreate(wordArr []rune, explain string) {
	ch := wordArr[0]
	var target *node
	for _, son := range n.sons {
		if son.word == ch {
			target = son
			break
		}
	}

	if target == nil {
		target = &node{word: ch}
		n.sons = append(n.sons, target)
	}

	if len(wordArr) == 1 {
		target.explain = explain
	} else {
		target.findOrCreate(wordArr[1:], explain)
	}
}

func (n *node) findSon(wordArr []rune) *node {
	ch := wordArr[0]
	var target *node
	for _, son := range n.sons {
		if son.word == ch {
			target = son
			break
		}
	}
	if target == nil || len(wordArr) == 1 {
		return target
	}

	return target.findSon(wordArr[1:])
}

func (n *node) query(word string) string {
	leaf := n.findSon([]rune(word))
	if leaf != nil {
		return leaf.explain
	}
	return ""
}

func (n *node) traverseAllLeafNodes() {
	s := &stack{}
	n.traverse(s)
}

func (n *node) traverse(s *stack) {

	if len(n.explain) > 0 {
		fmt.Println(n.explain)
	}

	if len(n.sons) == 0 {
		if next := s.pop(); next != nil {
			next.traverse(s)
		}
	} else {
		for i := 0; i < len(n.sons); i++ {
			son := n.sons[i]
			if i == (len(n.sons) - 1) {
				son.traverse(s)
			} else {
				s.push(son)
			}
		}
	}
}
