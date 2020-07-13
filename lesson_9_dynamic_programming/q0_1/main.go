package main

import "math"

type node struct {
	status int
}

func main() {
	source := []rune("添加一个字符")
	target := []rune("替换一个字符")

	// 1.初始化状态图
	matrix := make([][]int, len(source)+1)
	for i := 0; i <= len(source); i++ {
		matrix[i] = make([]int, len(target)+1)
	}
	// 2. 初始化边界状态
	for i := 0; i <= len(source); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(target); j++ {
		matrix[0][j] = j
	}

	// 3. 迭代状态
	for i := 1; i <= len(source); i++ {
		for j := 1; j <= len(target); j++ {
			add := matrix[i-1][j] + 1
			delete := matrix[i][j-1] + 1
			val := 1
			if source[i-1] == target[j-1] {
				val = 0
			}
			replace := matrix[i-1][j-1] + val
			matrix[i][j] = min(add, delete, replace)
		}
	}

	// 4. 获取结果
	println(matrix[len(source)][len(target)])
}

func min(a, b, c int) int {
	mid := math.Min(float64(a), float64(b))
	return int(math.Min(float64(c), mid))
}
