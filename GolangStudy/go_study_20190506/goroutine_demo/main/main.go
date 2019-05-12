package main

import (
	"fmt"
	"sync"
	"time"
)

//需求，使用goroutinue计算1-200的阶乘结果并保存到map中


var(
	myMap=make(map[int]int,10)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	lock sync.Mutex
)

func test(n int) {
	res:=1
	for i := 1; i <= n; i++ {
		res*=i
	}

	//加锁
	lock.Lock()
	//将res放入到map
	myMap[n]=res
	//解锁
	lock.Unlock()

}

func main() {
	//开启多个协程完成这个任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	time.Sleep(5*time.Second)

	lock.Lock()
	for i,v:=range myMap{
		fmt.Printf("myMap[%v]=%v\n",i,v)
	}
	lock.Unlock()
}
