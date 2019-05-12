package main

import "fmt"

//write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++{
		//放入数据
		intChan<-i
		fmt.Printf("writeData 写入数据=%v\n",i)
	}
	close(intChan)//关闭
}

func readData(intChan chan int, exitChan chan bool) {
	for{
		v,ok:=<-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n",v)
	}
	//readData 读取数据完成，任务完成
	exitChan<-true
	//带缓存的信道是必须写完后close，否则读取就会一直等待造成deadlock
	close(exitChan)
}

func main() {
	//创建两个管道
	inchan :=make(chan int,50)
	exitChan:=make(chan bool,1)

	go writeData(inchan)
	go readData(inchan,exitChan)

	for{
		_,ok:=<-exitChan
		if !ok {
			break
		}
	}
}
