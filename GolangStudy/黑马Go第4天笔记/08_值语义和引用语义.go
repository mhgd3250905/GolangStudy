package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值
//参数为普通变量，非指针，值语义，即为一份拷贝
func (p Person) SetInfoValue(n string,s byte,a int)  {
	p.name=n
	p.sex=s
	p.age=a
	fmt.Println("SetInfoValue &p = ",&p)
	fmt.Printf("SetInfoValue &p = %p\n",&p)

}

//接收者为指针变量
func (p *Person) SetInfoPointer(n string,s byte,a int)  {
	p.name=n
	p.sex=s
	p.age=a

	fmt.Println("SetInfoPointer p = ",p)
	fmt.Printf("SetInfoPointer p = %p\n",p)



}

func main() {
	var s1 Person=Person{"go",'m',22}
	fmt.Printf("&s1 = %p\n",&s1)
	//值语义
	s1.SetInfoValue("mike",'m',18)
	fmt.Println("s1 = ",s1)

	s1.SetInfoPointer("mike",'m',18)
	fmt.Println("s1 = ",s1)


}