package main

import (
	"fmt"
)

func main() {
	//定义一个数组 [10]int 和 [5]int 是不同类型
	//[数字]，这个数字是作为数组元素个数
	var arr [10]int

	var brr [5]int

	fmt.Printf("len(arr)=%d,len(brr)=%d", len(arr), len(brr))

	//注意：指定的数组元素个数必须是常量

	//操作数组元素，从0开始，到len()-1，部队称原则,这个数字，叫下标
	//下标可以是变量或常量
	arr[0] = 1
	i := 1
	arr[i] = 2 //a[1]=2
}
