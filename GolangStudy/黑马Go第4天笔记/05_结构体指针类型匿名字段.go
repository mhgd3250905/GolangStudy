package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*Person	//指针类型
	int	//基础类型的匿名字段
	addr string
	name string//和person同名了
}

func main() {
	s1:=Student{&Person{"mike",'m',18},666,"bj","yoyo"}
	fmt.Printf("s1 = %+v\n",s1)

	fmt.Println(s1.name,s1.sex,s1.age,s1.int,s1.addr)

	//先定义变量
	var s2 Student
	s2.Person=new(Person) //分配空间
	s2.name="yoyo"
	s2.sex='m'
	s2.age=18
	s2.int=666
	s2.addr="bj"
	fmt.Println(s2.name,s2.sex,s2.age,s2.int,s2.addr)

}