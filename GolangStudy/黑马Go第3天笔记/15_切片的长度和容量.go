package main

import (
	"fmt"
)

func main() {
	a:=[]int{1,2,3,0,0}
	s:=a[0:3:5]
	fmt.Println("s= ",s)
	fmt.Println("len(s)= ",len(s))
	fmt.Println("cap(s)= ",cap(s))

	fmt.Println()

	s=a[1:4:5]
	fmt.Println("s= ",s)//从下标1开始，取4-1=3个
	fmt.Println("len(s)= ",len(s))//长度 4-1=3个
	fmt.Println("cap(s)= ",cap(s))//容量 5-1=4个

}
