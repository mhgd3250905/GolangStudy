package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

//person类型，实现了一个方法
func (tmp *Person) PrintInfo(){
	fmt.Printf("name = %s,sex = %c,age = %d",tmp.name,tmp.sex,tmp.age)
}

//有一个学生继承了Person，成员和方法都继承了
type Student struct {
	Person//匿名字段
	id int
	addr string
}

//student也实现了一个方法，这个方法和Person方法重名
func (tmp *Student) PrintInfo(){
	fmt.Println("tmp = ",tmp)
}


func main() {
	s:=Student{Person{"mike",'m',18},666,"bj"}
	//就近·原则，先找本作用域的方法，找不到再用继承的方法
	s.PrintInfo()

	//显示调用继承的方法
	s.Person.PrintInfo()
}