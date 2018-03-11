package main

import(
	"fmt"
)

func main() {
	//a:=10//整形变量a

	var p *int
	//指向一个合法内存
	//p=&a

	//p是*int,指向int类型
	p=new(int)

	*p=666
	fmt.Printf("*p=%v\n",*p)

	q:=new(int)//自动推导类型
	*q=777
	fmt.Printf("*q type is %T",*q)
}
