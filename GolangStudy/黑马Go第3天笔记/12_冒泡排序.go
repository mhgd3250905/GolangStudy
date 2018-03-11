package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	//设置种子，只需要一次
	rand.Seed(time.Now().UnixNano()) //以当前系统时间作为种子参数

	var a [10]int
	n := len(a)

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100) //100以内的随机数
	}
	fmt.Println("a: ",a)

	fmt.Println("排序后")
	//冒泡排序，挨着的2个元素比较，升序（大于则交换）
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j]>a[j+1] {
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	fmt.Println("a: ",a)
}
