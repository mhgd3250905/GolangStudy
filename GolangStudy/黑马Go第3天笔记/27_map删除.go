package main

import (
	"fmt"
)

func main() {
	m1:=map[int]string{1:"mike",2:"yoyo",3:"go"}
	fmt.Println("m1 = ",m1)

	delete(m1,1)//删除key值为1的内容
	fmt.Println("m1 = ",m1)

}
