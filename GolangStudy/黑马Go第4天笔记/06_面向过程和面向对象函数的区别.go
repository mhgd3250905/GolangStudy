package main

import "fmt"

//实现两个数相加
//面向过程
func Add(a, b int) int {
	return a+b
}

//面向对象，方法：给某个类型绑定一个函数

type long int

func (tmp long) Add02(other long) long{
	return tmp+other
}


func main() {
	var result int
	result=Add(1,1)
	fmt.Println("result = ",result)

	//定义一个变量
	var a long=2
	//调用方法格式： 变量名.函数(参数)
	result2:=a.Add02(3)
	fmt.Println("result = ",result2)

	//面向对象只是换了一种表现形式

}
