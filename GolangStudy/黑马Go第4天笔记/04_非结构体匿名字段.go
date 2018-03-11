package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person	//只有类型，没有名字，你名字短，继承了Person成员
	int	//基础类型的匿名字段
	addr string
	name string//和person同名了
}

func main() {
	s:=Student{Person{"mike",'m',18},666,"bj","hehe"}
	fmt.Printf("s = %+v\n",s)

	s.int=777

	fmt.Println(s.name,s.addr,s.sex,s.int)
}