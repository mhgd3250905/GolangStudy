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
	s := Student{1, "mike", 'm', 18, "bj"}
	test01(&s) //地址传递（引用传递）,形参可以改实参
	fmt.Println("s = ",s)

}

func test01(s *Student) {
	s.id = 666
	fmt.Println("s = ",s)

}
