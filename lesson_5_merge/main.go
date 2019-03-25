/**
用递归实现归并排序
*/

package main

import "fmt"

func main() {
	fmt.Printf("%v\n", mergeSort([]int{1, 11, 5, 3, 6, 2, 4, 26, 99, 7, 27, 9, 8}))
}

func mergeSort(src []int) []int {

	if len(src) == 1 {
		return src
	}

	sorted1 := mergeSort(src[0 : len(src)/2])
	sorted2 := mergeSort(src[len(src)/2 : len(src)])
	sortedTotal := []int{}

	index1 := 0
	index2 := 0

	for len(sorted1) > index1 && len(sorted2) > index2 {
		item1 := sorted1[index1]
		item2 := sorted2[index2]
		if item1 < item2 {
			sortedTotal = append(sortedTotal, item1)
			index1++
		} else {
			sortedTotal = append(sortedTotal, item2)
			index2++
		}
	}

	if len(sorted1) > index1 {
		sortedTotal = append(sortedTotal, sorted1[index1:]...)
	} else {
		sortedTotal = append(sortedTotal, sorted2[index2:]...)
	}

	return sortedTotal
}
