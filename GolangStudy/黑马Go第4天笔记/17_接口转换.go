package main

import "fmt"

//定义接口类型
type Humaner interface {
	//子集
	//方法，只有声明没有实现,由别的类型实现
	sayHi()
}

type Personer interface {
	//超集
	Humaner //匿名字段，继承了sayHi()
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) sayHi() {
	fmt.Printf("Student[%s,%d] say hi!\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student在唱着： ", lrc)
}

func main() {
	//超集可以转换为子集，反过来不可以
	var iPro Personer //超集
	iPro=&Student{"mike",666}

	var i Humaner     //子集

	i=iPro
	i.sayHi()
}
