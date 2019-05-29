package main

import "fmt"

//定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //这个表示指向下一个结点
}

//给链表插入一个结点
//编写第一种插入方式，在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1.先找到该链表的最后一个结点
	//2.创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { //表示找到最后了
			break
		}
		temp = temp.next //让temp不断地指向下一个结点
	}
	//3.将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head

	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明链表中已经存在no了，就不让插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起已经存在这个no了。")
		return
	} else {
		//此时temp 为An temp.next 为A(n+1) 插入这个新的元素再次之间
		//所以让newNode.next指向A(n+1)再让An指向newNode
		newHeroNode.next=temp.next
		temp.next = newHeroNode
	}
}

func DelHeroNode(head *HeroNode, id int)  {
	temp := head

	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			//说明链表中已经存在no了，就不让插入
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		//找到删除
		temp.next=temp.next.next
	}else {
		fmt.Println("sorry 要删除的id不存在")
	}
}

//显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {
	//1.创建一个辅助结点
	temp := head

	//先判断该链表是否为一个空的链表
	if temp.next == nil {
		fmt.Println("该链表为空")
		return
	}

	//2.遍历这个链表
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		temp = temp.next
		//判断是否到队尾
		if temp.next == nil { //说明已经到最后一个结点了
			break
		}
	}
}

func main() {
	//1.先创建一个头结点
	head := &HeroNode{}

	//2.创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero3 := &HeroNode{
		no:      0,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero4 := &HeroNode{
		no:      4,
		name:     "宋江",
		nickname: "及时雨",
	}

	//3.加入
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero4)
	DelHeroNode(head, 4)
	//4.显示
	ListHeroNode(head)

}
