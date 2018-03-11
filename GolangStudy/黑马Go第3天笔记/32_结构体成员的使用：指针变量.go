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
	//指针有合法指向后，才操作成员
	//先定义一个普通结构体变量
	var s Student
	//定义一个指针变量，保存s的地址
	var p1 *Student
	p1=&s
	//通过指针操作 p1.id 和 (*p1),id 完全等价，只能使用 . 运算符
	p1.id=18
	(*p1).name="mike"
	p1.sex='m'
	p1.age=18
	p1.addr="bj"
	fmt.Println(" p1 = ",p1)

	//通过new 申请一个结构体
	p2:=new(Student)
	p2.id=18
	p2.name="mike"
	p2.sex='m'
	p2.age=18
	p2.addr="bj"
	fmt.Println(" p2 = ",p2)
}
