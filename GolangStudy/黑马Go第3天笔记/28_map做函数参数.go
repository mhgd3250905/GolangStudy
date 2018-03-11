package main

import (
	"fmt"
)

func main() {
	m1:=map[int]string{1:"mike",2:"yoyo",3:"go"}
	fmt.Println("m1 = ",m1)

	test(m1)//函数内部删除某个key
	fmt.Println("m1 = ",m1)

}
func test(m map[int]string) {
	delete(m,1)
}
