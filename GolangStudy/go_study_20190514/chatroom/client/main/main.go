package main

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/client/process"
	"fmt"
)

//定义用户id，用户密码
var userId int
var userPwd string
var userName string

func main() {
	//接收用户选择
	var key int

	for {
		fmt.Println("-----------------欢迎登陆多人聊天室-----------------")
		fmt.Println("                 1.登陆聊天室")
		fmt.Println("                 2.注册用户")
		fmt.Println("                 3.退出系统")
		fmt.Println("                 请选择（1-3）：")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id：")
			fmt.Scanf("%d\n", &userId) //如果不写回车就会把回车当做下一个输入记录到密码中
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录
			//1.创建一个UserProcess实例
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("up.Login fail err=", err)
			}
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%v\n",&userPwd)
			fmt.Println("请输入用户名称（nikename）:")
			fmt.Scanf("%v\n",&userName)
			//2,调用UserProcess,完成注册的请求
			up := &process.UserProcess{}
			err := up.Register(userId, userPwd,userName)
			if err != nil {
				fmt.Println("up.Login fail err=", err)
			}
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	////根据用户的输入，显示新的提示信息
	//if key == 1 {
	//	//用户登陆
	//	fmt.Println("请输入用户的id：")
	//	fmt.Scanf("%d\n", &userId) //如果不写回车就会把回车当做下一个输入记录到密码中
	//	fmt.Println("请输入用户密码：")
	//	fmt.Scanf("%s\n", &userPwd)
	//
	//	//创建
	//	//调用userProcess的登录
	//	//先把登陆函数写到另外一个文件
	//	//login(userId, userPwd)
	//
	//} else if key == 2 {
	//	fmt.Println("进行用户注册")
	//} else if key == 3 {
	//	fmt.Println("退出登陆")
	//	os.Exit(1)
	//}
}
