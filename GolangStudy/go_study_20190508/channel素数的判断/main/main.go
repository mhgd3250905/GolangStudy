package main

import "fmt"

func putNum(intChan chan int) {
	for i := 1; i <= 200000; i++ {
		intChan <- i
	}
	//关闭信道
	close(intChan)
}

//开启四个线程，从intChan中取出数据，并判断是否为素数
//如果是就放入到primeChan中
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool) {
	for{
		num,ok:=<-intChan
		if !ok {
			//如果已经取不到数字了，说明已经取完了，退出
			break
		}
		flag:=true
		for i := 2; i < num; i++ {
			//除了1和自己之外，还能够被其他数字整除，就不是素数
			if num%i==0 {
				flag=false
				break
			}
		}

		if flag {
			//如果是素数，就放入到primeChan中
			primeChan<-num
		}
	}
	fmt.Println("有一个协程因为取不到数据，退出了")
	//这里我们还不能关闭这个通道
	//向退出exitChna写入标识
	exitChan<-true
}

func main() {
	inChan := make(chan int, 1000)
	//放入结果
	primeChan := make(chan int, 2000)
	//开四个协程并行执行任务
	exitChan := make(chan bool, 4)

	//开启协程，向intChan放入数据 1-8000
	go putNum(inChan)

	go func() {
		//开启四个线程，从intChan中取出数据，并判断是否为素数
		//如果是就放入到primeChan中
		for i := 0; i < 4; i++ {
			go primeNum(inChan,primeChan,exitChan)
		}

		//主线程进行处理
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当我们从exitChan中取出了四个结果就可以关闭exitChan了
		close(exitChan)
		close(primeChan)
	}()

	//遍历primeChan 把结果取出
	for{
		res,ok:=<-primeChan
		if !ok {
			break
		}
		//输出结果
		fmt.Printf("素数： %v\n",res)
	}

	fmt.Println("main线程退出")
}
