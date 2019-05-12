package main

import (
	"fmt"
	"runtime"
)

func main() {
	//查看有多少个CPU
	cpuNum:=runtime.NumCPU()
	fmt.Println("cpuNum=",cpuNum)

	//自己设置使用多少个CPU
	runtime.GOMAXPROCS(cpuNum-1)
	fmt.Println("设置cpu完毕！")
}
