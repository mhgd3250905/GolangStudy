package main

import "fmt"

//定义接口类型
type Humaner interface {
	//方法，只有声明没有实现,由别的类型实现
	sayHi()
}

type Personer interface {
	Humaner//匿名字段，继承了sayHi()
	sing(lrc string)
}

type Student struct {
	name string
	id int
}

func (tmp *Student) sayHi(){
	fmt.Printf("Student[%s,%d] say hi!\n",tmp.name,tmp.id)
}

func (tmp *Student) sing(lrc string)  {
	fmt.Println("Student在唱着： ",lrc)
}

func main() {
	//定义一个接口类型的变量
	var i Personer
	s:=&Student{"mike",666}
	i=s
	i.sayHi()
	i.sing("爱的供养")
}