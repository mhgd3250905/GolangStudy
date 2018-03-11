package main

import (
	"fmt"
)

func main() {
	//自动推导类型,同时初始化
	s1:=[]int{1,2,3,4}
	fmt.Println("s1 = ",s1)

	//借助make函数,格式 make(slice,len,cap)
	s2:=make([]int,5,10)
	fmt.Printf("len= %d ,cap= %d \n",len(s2),cap(s2))

	//没有指定容量，容量和长度一样
	s3:=make([]int,5)
	fmt.Printf("len= %d ,cap= %d \n",len(s3),cap(s3))

}

func main1() {


	//切片和数组的区别
	//数组[]里面的长度是固定的一个常量，数组是不能修改长度，len和cap永远都是5
	a:=[5]int{}
	fmt.Printf("len= %d ,cap= %d ",len(a),cap(a))

	fmt.Println()

	//切片，[]里面为空或者为...，切片的长度或容量可以不固定
	s:=[]int{}
	fmt.Printf("len= %d ,cap= %d ",len(s),cap(s))

	s=append(s,11)//给切片末尾追加一个成员

	fmt.Printf("append： len= %d ,cap= %d ",len(s),cap(s))

}
