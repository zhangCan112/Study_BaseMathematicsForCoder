/**
实现一个密码暴力破解Demo,不用递归实现,实现最大长度的破解.
这里复习用到了位运算和进制的概念
*/
package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var inputPw = flag.String("pw", "a", "待测试密码")

var pwSet = [5]string{"a", "b", "c", "d", "e"}

type compareFunc = func(src string) bool

func main() {
	flag.Parse()
	compare := createCompare(*inputPw)
	result, _ := getPassword(compare, len(*inputPw))

	fmt.Println(result)
}

func getPassword(compare compareFunc, maxLength int) (string, bool) {
	for length := 1; length <= maxLength; length++ {
		kindsFloat := math.Pow(float64(len(pwSet)), float64(length))
		kinds := int(kindsFloat)
		base := len(pwSet)
		for kind := 0; kind < kinds; kind++ {
			//10进制转为pwSet.length的进制
			numStr := strconv.FormatInt(int64(kind), base)
			var srcPw = ""
			for _, indexStr := range []byte(numStr) {
				index, ok := strconv.ParseInt(string(indexStr), base, 64)
				if ok == nil {
					srcPw = pwSet[index] + srcPw
				}
			}
			if compare(srcPw) == true {
				return srcPw, true
			}
		}
	}

	return "未找到正确密码", false
}

///返回比较函数
func createCompare(target string) compareFunc {
	return func(src string) bool {
		return strings.Compare(target, src) == 0
	}
}
