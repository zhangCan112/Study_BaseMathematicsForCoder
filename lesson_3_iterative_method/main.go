/**
用迭代法实现求一个数的平方根
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(squareRoot(99, 0.001))
}

func squareRoot(num float64, precision float64) float64 {

	var start float64
	var end = num
	var result = (end + start) / 2.0
	for result > (1 + precision) {
		fmt.Println(result)
		if (math.Abs((result * result) - num)) < precision {
			return result
		} else if (result * result) > num {
			end = result
		} else {
			start = result
		}
		result = (end + start) / 2.0
	}

	return result
}
