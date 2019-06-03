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
func DelCatNode(head *CatNode, id int) *CatNode {

	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，无法删除")
		return head
	}

	//如果只有一个结点
	if temp.next == head { //只有一个结点
		temp.next = nil
		return head
	}

	//将helper定位到环形列表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果有两个以上的结点
	flag := true
	for {
		if temp.next == head {
			//如果已经比较到最后一个，且最后一个还没有比较
			break
		}
		if temp.no == id {
			if temp == head { //说明删除的的head
				head = head.next
			}
			//恭喜找到,我们也可以直接删除
			helper.next = temp.next
			flag = false
			break
		}
		temp = temp.next     //移动【比较】
		helper = helper.next //移动 【一旦找到要删除的结点 helper】
	}

	//这里还要比较一次
	if flag { //如果flag为真 我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Println("cat =%d\n",id)
		}else{
			fmt.Println("对不起没有找到")
		}
	}

	return head

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
	head=DelCatNode(head,2)
	head=DelCatNode(head,20)
	ListCircleLink(head)

}
