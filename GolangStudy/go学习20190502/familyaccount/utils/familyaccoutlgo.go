package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//声明一个变量，保存接受用户输入的选项
	key string
	//声明一个变量控制是否退出for循环
	loop bool
	//定义账户的余额
	balance float64
	//每次收支的今儿
	money float64
	//每次收支的说明
	note string
	//定义一个变量，记录是否有收支行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details进行拼接处理
	details string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:"",
		loop:true,
		balance:10000.0,
		note:"",
		flag:false,
		details:"收支\t\t账户金额\t\t收支净额\t\t说明",
	}
}



//讲显示明细写成一个方法
func (this *FamilyAccount) showDetails(){
	fmt.Println("\n--------------当前收支明细记录--------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧！")
	}
}

//给该结构体绑定方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {
	//显示主菜单
	for {
		fmt.Println("--------------家庭收支记账软件--------------")
		fmt.Println("                1 收支明细")
		fmt.Println("                2 登记收入")
		fmt.Println("                3 登记支出")
		fmt.Println("                4 退出软件")
		fmt.Println("请选择（1-4）:")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}

		if !this.loop {
			break
		}
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details
	this.details += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	fmt.Println(this.details)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("登记支出..")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	fmt.Println(this.details)
	this.flag = true
}

func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}