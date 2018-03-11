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
	var s Student

	//操作成员，需要时用 . 运算符
	s.id=1
	s.name="mike"
	s.sex='m'
	s.age=18
	s.addr="bj"

	fmt.Println("s = ",s)
}
