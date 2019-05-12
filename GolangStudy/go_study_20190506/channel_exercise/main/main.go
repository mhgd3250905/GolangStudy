package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main() {
	//定义一个可以存放任意类型的管道 3个数据
	allChan:=make(chan interface{},3)

	allChan<-10
	allChan<-"Tom Jack"
	cat:=Cat{"小花猫",4}
	allChan<-cat

	//我们希望获取到管道中的第三个元素，则先将前两个推出
	<-allChan
	<-allChan
	newCat:=<-allChan//从管道中取出的cat是什么
	fmt.Printf("newCat=%T,newCat=%v\n",newCat,newCat)
	a:=newCat.(Cat)
	fmt.Printf("newCat.Name=%v",a.Name)
}
