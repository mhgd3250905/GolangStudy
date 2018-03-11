package main

import (
	"fmt"
)

func main() {
	m1:=map[int]string{1:"mike",2:"yoyo"}
	//赋值如果村存在就覆盖
	fmt.Println("m1 = ",m1)
	m1[1]="c++"
	fmt.Println("m1 = ",m1)

}
