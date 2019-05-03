package main

import (
	"GolangStudy/GolangStudy/go学习20190502/CustomManager/modle"
	"GolangStudy/GolangStudy/go学习20190502/CustomManager/service"
	"fmt"
)

type customerView struct {
	//定义必要的字段
	key  string //接受用户的输入
	loop bool   //是否循环显示菜单
	//增加customerService字段
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {
	//获取到当前所有的客户信息
	customers := this.customerService.List()
	fmt.Println("------------客  户  列  表------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//输出客户信息
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------------客户列表完成------------")
}

//得到用户的输入信息，构建新的客户，并完成添加
func (this *customerView) Add() {
	fmt.Println("------------添  加  客  户------------")
	fmt.Println("姓名：")
	name:=""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender:=""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age:=0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone:=""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email:=""
	fmt.Scanln(&email)

	//构建一个新的Customer
	//注意Id 系统分配 唯一
	customer:=modle.NewCustomerWithoutId(name,gender,age,phone,email)

	if this.customerService.Add(customer) {
		fmt.Println("------------添  加  完  成------------")
	}else {
		fmt.Println("------------添  加  失  败------------")
	}
}

//得到用户的输入，删除该Id对应的客户
func (this *customerView) delete() {
	fmt.Println("------------删  除  客  户------------")
	fmt.Println("请输入待删除的客户编号（-1退出）：")
	id:=-1
	fmt.Scanln(&id)
	if id == -1 {
		return	//放弃删除
	}

	fmt.Println("确认是否删除（Y/N）：")
	choice:=""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" ||choice == "Y" || choice == "N"{
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "Y" || choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("------------删  除  完  成------------")
		}else {
			fmt.Println("------------删  除  失  败------------")
		}
	}else {
		return
	}
}

//得到用户的输入,更新用户的数据
func (this *customerView) update() {
	fmt.Println("------------修  改  客  户------------")
	fmt.Println("请输入待修改的客户编号（-1退出）：")
	id:=-1
	fmt.Scanln(&id)
	if id == -1 {
		return	//放弃删除
	}

	index:=this.customerService.FindById(id)
	if index == -1 {
		//没有这个客户
		fmt.Printf("无法找到编号为%v的客户...",id)
		return
	}

	fmt.Println("姓名：")
	name:=""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender:=""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age:=0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone:=""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email:=""
	fmt.Scanln(&email)
	fmt.Println("确认是否修改（Y/N）：")
	choice:=""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" ||choice == "Y" || choice == "N"{
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	if choice == "Y" || choice == "y" {
		if this.customerService.Update(id,modle.NewCustomerWithoutId(name,gender,age,phone,email)) {
			fmt.Println("------------修  改  完  成------------")
		}else {
			fmt.Println("------------修  改  失  败------------")
		}
	}else {
		return
	}
}


//退出软件
func (this *customerView)Exit()  {
	fmt.Println("确认是否退出（Y/N）：")
	for{
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入...")
	}

	if this.key=="Y"||this.key=="y" {
		this.loop=false
	}
}

//显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("\n------------客户信息管理软件------------")
		fmt.Println("             1 添 加 客 户")
		fmt.Println("             2 修 改 客 户")
		fmt.Println("             3 删 除 客 户")
		fmt.Println("             4 客 户 列 表")
		fmt.Println("             5 退 出")
		fmt.Println("请选择（1-5）：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.Add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.Exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")

		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")

}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//完成对CustomerService的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
