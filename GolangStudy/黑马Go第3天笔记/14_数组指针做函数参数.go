package main

import (
	"fmt"
)

func main() {
	a:=[5]int{1,2,3,4,5}//初始化

	modify(&a)//地址传递
	fmt.Println("main:a= ",a)
}

//p指向实参数组a,它是指向数组，它是数组指针
//*p代表指针所指向的内存，就是实参
func modify(p *[5]int) {
	(*p)[0]=666
	fmt.Println("modify:a= ",*p)

}
