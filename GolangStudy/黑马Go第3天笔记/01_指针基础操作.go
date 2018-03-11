package main

import(
	"fmt"
)

func main() {
	var a int =10
	//每个变量有2层含义：变量的内存，变量的地址
	fmt.Printf("a=%d\n",a)
	fmt.Printf("&a=%v\n",&a)

	//保存某个变量的地址，需要指针类型  *int 保存int的地址 **int  保存*int 地址
	//声明（定义），定义知识特殊的声明
	var p *int
	p=&a//指针变量指向谁，就把谁的地址复制给指针变量
	fmt.Printf("p=%v,&a=%v\n",p,&a)

	*p=666//*p操作的不是p的内存，是p所指向的内存（也就是a）
	fmt.Printf("*p=%v,a=%v\n",*p,a)
}
