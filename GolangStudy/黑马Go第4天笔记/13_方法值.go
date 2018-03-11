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

	s.SetInfoPointer()//传统调用方式

	//保存函数入口地址
	pFunc:=s.SetInfoPointer//方法值，调用函数时无需再传递接收者，隐藏接受者
	pFunc()

}