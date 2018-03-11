package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person	//只有类型，没有名字，你名字短，继承了Person成员
	id   int
	addr string
}

func main() {
	var s1 Student=Student{Person{"mike",'m',18},1,"bj"}
	fmt.Println(s1.name,s1.sex,s1.age,s1.addr)
	s1.name="yoyo"
	s1.sex='f'
	s1.age=22
	s1.id=666
	s1.addr="sz"
	fmt.Println(s1.name,s1.sex,s1.age,s1.addr)

}
