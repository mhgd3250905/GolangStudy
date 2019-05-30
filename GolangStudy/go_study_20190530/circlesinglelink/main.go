package main

import "fmt"

type CatNode struct {
	no   int //猫猫的编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //形成一个环状
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	//先定义一个临时的变量，帮忙找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	newCatNode.next = head
	temp.next = newCatNode

}

//删除一只猫
func DelCatNode(head *CatNode, id int) {

}

//输出这个环形的链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下:")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}

	for {
		fmt.Printf("猫的信息为= [id=%d name=%s]->", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {

	//这里我么你初始化一个环形链表的头结点
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "jerry",
	}
	cat3 := &CatNode{
		no:   3,
		name: "white",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

}
