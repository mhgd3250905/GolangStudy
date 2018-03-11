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
	fmt.Println("s1 = ",s1)

	s2:=Student{Person{"mike",'m',18},1,"bj"}
	fmt.Println("s2 = ",s2)
	fmt.Printf("s2 =%+v\n",s2)

	//指定成员初始化，没有初始化测成员自动赋值为0
	s3:=Student{id:1}
	fmt.Printf("s3 = %+v",s3)
}
