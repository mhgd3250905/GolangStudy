package main

import (
	"fmt"
)

func main() {
	//定义一个变量，类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len，没有cap
	fmt.Printf("len = %d\n",len(m1))

	m2:=make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Printf("len = %d\n",len(m2))

	//可以通过make创建，可以指定长度，只是指定了容量，但是里面确实一个数据也没有
	m3:=make(map[int]string,10)
	fmt.Println("m3 = ",m3)
	fmt.Printf("len = %d\n",len(m3))

	for i := 0; i < 20; i++ {
		m3[i]=fmt.Sprintf("value %d",i)
	}

	//键值对是唯一的
	fmt.Println("m3 = ",m3)

}
