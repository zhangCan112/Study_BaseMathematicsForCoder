/*
动态规划
计算任意两个字符串的最优编辑距离，并给出最优解的执行方式
*/

package main

import (
	"fmt"
	"math"
)

///可用策略
const (
	ADD     = iota //添加一个字符
	DELETE         //删除一个字符
	REPLACE        //替换一个字符
)

func main() {
	var editDistance = dynamicProgramming("添加一个字符", "替换一个字符")
	fmt.Println(editDistance)
}

func dynamicProgramming(src string, target string) int {
	srcArr := []rune(src)
	targetArr := []rune(target)

	//初始化一个二位数组
	matirx := make([][]int, len(srcArr)+1)
	for i := 0; i <= len(srcArr); i++ {
		matirx[i] = make([]int, len(targetArr)+1)
	}

	//初始化数据
	for i := 0; i <= len(srcArr); i++ {
		matirx[i][0] = i
	}
	for j := 0; j <= len(targetArr); j++ {
		matirx[0][j] = j
	}

	for i := 1; i <= len(srcArr); i++ {
		for j := 1; j <= len(targetArr); j++ {
			first := matirx[i-1][j] + 1
			second := matirx[i][j-1] + 1
			flag := 1
			if srcArr[i-1] == targetArr[j-1] {
				flag = 0
			}
			third := matirx[i-1][j-1] + flag
			matirx[i][j] = min(first, second, third)
		}
	}

	fmt.Println(matirx)
	return matirx[len(srcArr)][len(targetArr)]
}

func min(first int, second int, third int) int {
	min := math.Min(float64(first), float64(second))
	min = math.Min(min, float64(third))
	return int(min)
}
