package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type CircleQueue struct {
	maxSize int
	array   [5]int //数组模拟队列
	head    int    //表示指向队列首
	tail    int    //表示指向队列尾
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("circlequeue full")
	}
	//this.tail在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val //把值给尾部
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("circlequeue empty")
	}
	//取出 head是指向队首，并且含队首元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

//判断环形队列满了
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判断环形队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出唤醒队列有多少个元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

//显示队列
//找到队首，然后遍历到队尾
func (this *CircleQueue) ListQueue() {
	//取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
}

func main() {
	//先创建一个队列
	queue := &CircleQueue{
		maxSize: 5,
	}

	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要加入的对列数：")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println("加入队列失败，err=", err)
			} else {
				fmt.Println("加入队列成功")
				fmt.Println()
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println("获取数据失败，err=", err)
			}
			fmt.Println("get= ", val)
		case "show":
			queue.ListQueue()
			fmt.Println()
		case "exit":
			os.Exit(0)
		}
	}
}
