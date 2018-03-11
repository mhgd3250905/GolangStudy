package main

import (
	"fmt"
)

func main() {
	m1:=map[int]string{1:"mike",2:"yoyo",3:"go"}
	//第一个返回值为key,第二个返回值为value，遍历结果为无序的
	for key,value:=range m1{
		fmt.Printf("%d ----> %s\n",key,value)
	}

	//如何判断一个key值是否存在
	//第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在ok为true
	value,ok:=m1[1]
	if ok {
		fmt.Println("m1[1] = ",value)
	}else {
		fmt.Println("key值不存在")
	}
}
