package main

import "fmt"

func main() {
	//管道可以生命为只读或者只写

	//1.默认情况下，管道是双向的
	//var chan1 chan int

	//2.声明为只写
	var chan2 chan<- int
	chan2=make(chan int)
	chan2<-20
	//num:=<-chan2//报错

	//3.声明为只读
	var chan3 <-chan int
	num2:=<-chan3
	//chan3<-30//报错
	fmt.Println(num2)


}

//作为参数 这个管道只做写入操作
func send(ch chan<- int){
	ch<-100
}

//作为参数这个管道只做读取操作
func recv(ch <-chan int){
	num,_:=<-ch
	fmt.Println(num)
}
