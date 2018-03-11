package main

import "fmt"

//定义一个结构体类型

type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	//顺序初始化，每个成员都必须初始化
	var p1 *Student=&Student{1,"mike",'m',18,"bj"}
	fmt.Println("*p1 = ",*p1)

	p2:=&Student{name:"mike",addr:"bj"}
	fmt.Println("*p2 = ",*p2)

}
