package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int //数组模拟队列
	front   int    //表示指向队列首
	rear    int    //表示指向队列尾
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断是否已满
	if this.rear==this.maxSize-1{//重要的提示，rear是队列尾部（含最后元素）
		return errors.New("circlequeue full")
	}

	this.rear++//rear后移
	this.array[this.rear]=val
	return
}

//显示队列
//找到队首，然后遍历到队尾
func (this *Queue)ShowQueue()  {
	//this.front不包含队首的元素
	for i:=this.front+1;i<=this.rear;i++{
		fmt.Printf("array[%d]=%d\n",i,this.array[i])
	}
}

//从队列中取出数据
func (this *Queue)GetQueue() (val int,err error) {
	//先判断队列是否为空
	if this.rear==this.front{
		return -1,errors.New("circlequeue empty")
	}
	this.front++
	val=this.array[this.front]
	return val,err
}

func main() {
	//先创建一个队列
	queue:=&Queue{
		maxSize:5,
		front:-1,
		rear:-1,
	}

	var key string
	var val int
	for{
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要加入的对列数：")
			fmt.Scanln(&val)
			err:=queue.AddQueue(val)
			if err != nil {
				fmt.Println("加入队列失败，err=",err)
			}else{
				fmt.Println("加入队列成功")
				fmt.Println()
			}
		case "get":
			val,err:=queue.GetQueue()
			if err != nil {
				fmt.Println("获取数据失败，err=",err)
			}
			fmt.Println("get= ",val)
		case "show":
			queue.ShowQueue()
			fmt.Println()
		case "exit":
			os.Exit(0)
		}
	}
}
