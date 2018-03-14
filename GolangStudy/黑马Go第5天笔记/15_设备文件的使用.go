package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Stdout.Close()//关闭后无法输出
	//fmt.Println("Are you OK!")
	//os.Stdout//默认已经打开，用户可以直接使用

	os.Stdout.WriteString("Are you OK")

	var a int
	fmt.Println("请输入a: ")
	fmt.Scan(&a)
	fmt.Println("a = ",a)

}
