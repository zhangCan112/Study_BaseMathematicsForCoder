package main

import "math"

func main() {
	// 1.初始化状态图

	// 2. 初始化边界状态

	// 3. 迭代状态

	// 4. 获取结果
}

func min(a, b, c int) int {
	mid := math.Min(float64(a), float64(b))
	return int(math.Min(float64(c), mid))
}
