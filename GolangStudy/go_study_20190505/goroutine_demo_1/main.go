package main

import (
	"fmt"
	"strconv"
	"time"
)

//在主线程（也可以理解为进程）中，开启一个goroutine，该协程每隔1秒输出"Hello World"
//在主线程中也每隔一秒输出"Hello golang" 输出10次后，退出程序
//要求主线程和goroutine同时进行

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test hello world "+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main()  {

	go test()


	for i := 0; i < 10; i++ {
		fmt.Println("main hello golang "+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
