package main

import (
	"fmt"
	"time"
)

func main() {
	//使用select可以解决读取数据的阻塞问题

	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道 5个数据string
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时候，如果不关闭阻塞而导致 deadlock
	//问题，在实际开发中，可能不好确定什么时候关闭管道
	//可以使用select 方式可以解决
	label:
	for {
		select {
		//注意，这里如果intChan一直没有关闭，不会一直阻塞而导致deadlock
		//会自动到写一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan中取出int:%v\n", v)
			time.Sleep(500*time.Millisecond)
		case v := <-strChan:
			fmt.Printf("从strChan中取出str:%v\n", v)
			time.Sleep(500*time.Millisecond)
		default:
			fmt.Printf("都取不到了，程序员可以加入逻辑")
			break label//或者使用return
		}
	}

}
