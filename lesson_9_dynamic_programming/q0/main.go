/*
动态规划
计算任意两个字符串的最优编辑距离，并给出最优解的执行方式
*/

package main

import "fmt"

///可用策略
const (
	ADD     = iota //添加一个字段
	DELETE         //删除一个字符
	REPLACE        //替换一个字符
)

func main() {
	dynamicProgramming("张灿1221", "123")
}

func dynamicProgramming(src string, target string) int {
	srcArr := []rune(src)
	targetArr := []rune(target)
	matirx := make([][]int, len(srcArr))
	for i := 0; i < len(srcArr); i++ {
		matirx = append(matirx, make([]int, len(targetArr)))
	}

	matirx[2][2] = 1123
	fmt.Println(matirx)
	return 1
}
