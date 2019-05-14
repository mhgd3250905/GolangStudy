package main

import (
	"fmt"
	"os"
)

//定义用户id，用户密码
var userId int
var userPwd string

func main() {
	//接收用户选择
	var key int
	//判断是否还继续显示菜单
	var loop = true

	for loop{
		fmt.Println("-----------------欢迎登陆多人聊天室-----------------")
		fmt.Println("                 1.登陆聊天室")
		fmt.Println("                 2.注册用户")
		fmt.Println("                 3.退出系统")
		fmt.Println("                 请选择（1-3）：")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop=false
		case 2:
			fmt.Println("注册用户")
			loop=false
		case 3:
			fmt.Println("退出系统")
			loop=false
		default:
			fmt.Println("你的输入有误，请重新输入")

		}
	}

	//根据用户的输入，显示新的提示信息
	if key == 1 {
		//用户登陆
		fmt.Println("请输入用户的id：")
		fmt.Scanf("%d\n",&userId)//如果不写回车就会把回车当做下一个输入记录到密码中
		fmt.Println("请输入用户密码：")
		fmt.Scanf("%s\n",&userPwd)
		//先把登陆函数写到另外一个文件
		err:=login(userId,userPwd)
		if err != nil {
			fmt.Println("登陆失败")
		}else {
			fmt.Println("登陆成功")
		}
	}else if key == 2 {
		fmt.Println("进行用户注册")
	}else if key ==3 {
		fmt.Println("退出登陆")
		os.Exit(1)
	}
}
