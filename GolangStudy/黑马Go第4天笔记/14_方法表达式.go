package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue()  {

fmt.Printf("SetInfoValue: %p, %v\n",&p,p)

}


func (p *Person) SetInfoPointer()  {
	fmt.Printf("SetInfoPointer: %p, %v\n",p,p)
}

func main() {
	s:=Person{"mike",'m',18}
	fmt.Printf("main: %p, %v\n",&s,s)

	//方法表达式
	f:=(*Person).SetInfoPointer
	f(&s)//显示把接受者传递过去

	f2:=(Person).SetInfoValue
	f2(s)
}