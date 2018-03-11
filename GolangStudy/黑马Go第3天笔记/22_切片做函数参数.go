package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int =10

	//创建一个切片，len为n
	s:=make([]int,n,n)

	initData(s)//初始化数据

	fmt.Println("排序前",s)

	BobbleSort(s)

	fmt.Println("排序后",s)


}

//初始化数据
func initData(s []int) {
	//设置随机种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i]=rand.Intn(100)//100以内的随机数
	}
}

//冒泡排序
func BobbleSort(s []int){
	n:=len(s)
	for i := 0; i<n-1;i++{
		for j := 0; j < n-1-i; j++ {
			if s[j]>s[j+1] {
				s[j],s[j+1]=s[j+1],s[j]
			}
		}
	}
}
