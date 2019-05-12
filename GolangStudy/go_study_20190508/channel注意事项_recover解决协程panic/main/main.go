package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world!")
	}
}

func test() {
	//这里可以使用defer+recover解决
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test()发生错误，",err)
		}
	}()

	//定义一个map
	var myMap map[int]string
	//这里没有对map进行初始化，所以会直接报错
	myMap[0]="golang"
}

func main() {

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Printf("main() ok=%v\n",i)
		time.Sleep(time.Second)
	}

}
