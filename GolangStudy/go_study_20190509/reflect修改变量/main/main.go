package main

import (
	"fmt"
	"reflect"
)

//通过反射修改 num int的值
//修改student的值

func reflect01(b interface{})  {
	//获取到reflect.Value
	rVal:=reflect.ValueOf(b)
	//看看rVal的kind 是一个指针
	fmt.Printf("rVal kind = %v\n",rVal.Kind())
	//设置值
	rVal.Elem().SetInt(20)
}

func main() {
	var num int =10
	reflect01(&num)
	fmt.Println("num = ",num)

	//可以这样理解rVal.Elem()
	//相当于取到了指针指向的地址空间的值
}