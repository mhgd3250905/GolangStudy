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
	name string//和person同名了
}

func main() {
	var s Student

	//默认规则，就近原则：
	//如果能在本作用域找到次成员，那么就操作此成员，如果没有找到，就找继承的字段
	s.name="mike"
	s.sex='m'
	s.age=18
	s.addr="bj"

	//显示调用
	s.Person.name="yoyo"

	fmt.Printf("s = %+v",s)
}
