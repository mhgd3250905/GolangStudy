package main

import "fmt"

//定义接口类型
type Humaner interface {
	//方法，只有声明没有实现,由别的类型实现
	sayHi()
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) sayHi(){
	fmt.Printf("Student[%s,%d] say hi!\n",tmp.name,tmp.id)
}

type Teacher struct {
	addr string
	group string
}

func (tmp *Teacher) sayHi(){
	fmt.Printf("Teacher[%s,%s] say hi!\n",tmp.addr,tmp.group)
}

type MyStr string

func (tmp *MyStr) sayHi(){
	fmt.Printf("MyStr[%s] say hi!\n",*tmp)
}

//定义了一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同表现，多态
func WhoSayHi(i Humaner){
	i.sayHi()
}

func main() {
	s:=&Student{"mike",66}
	t:=&Teacher{"bj","go"}
	var str MyStr ="Hello mike"

	//调用同一函数，就有不同表现
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x:=make([]Humaner,3)
	x[0]=s
	x[1]=t
	x[2]=&str

	//第一个返回下标，第二个返回下标对应的值
	for _,i:=range x{
		i.sayHi()
	}
}


func main01() {
	//定义接口类型的变量
	var i Humaner

	//只要是实现了该接口的方法的类型，那么这个类型的变量（接受者类型）就可以给i赋值
	s:=&Student{"mike",66}
	i=s
	i.sayHi()

	t:=&Teacher{"bj","go"}
	i=t
	i.sayHi()

	var str MyStr ="Hello mike"
	i=&str
	i.sayHi()
}