package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue()  {
	fmt.Println("SetInfoValue")

}


func (p *Person) SetInfoPointer()  {
	fmt.Println("SetInfoPointer")



}

func main() {
	//假如结构体变量是一个指针变,量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p:=Person{"mike",'m',18}
	p.SetInfoPointer()//内部先把p转化为&p，再调用


}