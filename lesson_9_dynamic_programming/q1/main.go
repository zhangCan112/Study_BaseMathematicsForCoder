/**
动态规划练习题。 有2，3，7三种钱币，求有最少的钱币凑满100元
分别使用 迭代，递归 和动态规划的思路解决
*/

package main

import (
	"fmt"
	"time"
)

//金额类别
var amountCategory = [3]int{2, 3, 7}

func main() {
	//递归法调用
	// s1 := time.Now().Unix()
	// fmt.Println(recursion(60))
	// e1 := time.Now().Unix()
	// fmt.Println((e1 - s1))

	//迭代法调用
	s2 := time.Now().Unix()
	// fmt.Println(iteration(10000111))
	iteration(1000011111)
	e2 := time.Now().Unix()
	fmt.Println((e2 - s2))

	//动态规划
	s3 := time.Now().Unix()
	// fmt.Println(optimalDynamicProgramming(10000111))
	optimalDynamicProgramming(1000011111)
	e3 := time.Now().Unix()
	fmt.Println((e3 - s3))
}

/*
迭代法
*/
func iteration(amount int) []int {

	for index := amount / 7; index >= 0; index-- {
		initial7 := initialArr(7, index)
		if len(initial7)*7 == amount {
			return initial7
		}

		leftAmount := amount - len(initial7)*7
		for index1 := leftAmount / 3; index1 >= 0; index1-- {
			initial3 := initialArr(3, index1)
			if len(initial3)*3 == leftAmount {
				return append(initial7, initial3...)
			}

			leftAmount2 := leftAmount - len(initial3)*3
			for index2 := leftAmount2 / 2; index2 >= 0; index2-- {
				initial2 := initialArr(2, index2)
				if len(initial2)*2 == leftAmount2 {
					return append(initial7, append(initial3, initial2...)...)
				}
			}
		}

	}
	return nil
}

func initialArr(defaultValue int, length int) []int {
	var origin = make([]int, length)
	for index := 0; index < length; index++ {
		origin[index] = defaultValue
	}

	return origin
}

/*
递归法
*/
func recursion(amount int) []int {

	if amount == 0 {
		return make([]int, 0)
	}

	if amount < 0 {
		return nil
	}

	first := recursion(amount - 2)
	if first != nil {
		first = append(first, 2)
	}

	second := recursion(amount - 3)
	if second != nil {
		second = append(second, 3)
	}

	third := recursion(amount - 7)
	if third != nil {
		third = append(third, 7)
	}

	min := minArr(first, second)
	min = minArr(min, third)
	return min
}

func minArr(a []int, b []int) []int {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if len(a) < len(b) {
		return a
	}
	return b
}

/*
动态规划法
*/
func optimalDynamicProgramming(total int) []int {

	for i := 0; i < total; i++ {
		left := total - i
		maxAmount := amountCategory[len(amountCategory)-1]
		if left%maxAmount == 0 {
			midResult := dynamicProgramming(i)
			if midResult != nil {
				return append(midResult, initialArr(maxAmount, left/maxAmount)...)
			}
		}
	}

	return nil
}

func dynamicProgramming(total int) []int {
	optimalArr := make([][]int, total+1)
	//初始条件
	optimalArr[0] = make([]int, 0)

	//优化初始条件
	maxAmount := amountCategory[len(amountCategory)-1]
	for index := 1; index <= total/maxAmount; index++ {
		optimalArr[index*maxAmount] = initialArr(maxAmount, index)
		//进一步添加尽可能多的初始条件
		for i := 0; i < len(amountCategory)-1; i++ {
			if (index*maxAmount+amountCategory[i]) <= total && optimalArr[index*maxAmount+amountCategory[i]] == nil {
				optimalArr[index*maxAmount+amountCategory[i]] = append(optimalArr[index*maxAmount], amountCategory[i])
			}
		}
	}

	//开始递推
	for i := 1; i < total+1; i++ {
		//初始条件已包含该最优解，则不用再计算
		if optimalArr[i] != nil {
			continue
		}
		var min []int
		for _, amount := range amountCategory {
			if amount > i {
				continue
			}
			left := i - amount
			for _, optimal := range optimalArr {
				if optimal != nil && sum(optimal) == left {
					min = minArr(min, append(optimal, amount))
				}
			}
		}
		optimalArr[i] = min

		if optimalArr[i] != nil && (i+amountCategory[len(amountCategory)-1] <= total) {
			optimalArr[i+amountCategory[len(amountCategory)-1]] = append(optimalArr[i], amountCategory[len(amountCategory)-1])
		}
	}

	return optimalArr[total]
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
