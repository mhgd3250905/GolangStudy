package process

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/client/utils"
	"fmt"
	"net"
	"os"
)

//显示登陆成功后的界面..
func ShowMenu() {
	for {
		fmt.Println("-----------------恭喜xxx登录成功-----------------")
		fmt.Println("                 1.显示在线用户列表")
		fmt.Println("                 2.发送消息")
		fmt.Println("                 3.信息列表")
		fmt.Println("                 4.退出系统")
		fmt.Println("请选择（1-4）：")
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择了退出系统...")
			os.Exit(0)
		default:
			fmt.Println("你输入的选择不正确...")
		}
	}
}

//和服务器端保持通讯
func serverProcessMsg(conn net.Conn) {
	//创建一个transfer实例，不停地读取服务器发送的消息
	tf:=&utils.Transfer{
		Conn:conn,
	}
	for{
		fmt.Println("客户端正在等待读取服务器发来的消息")
		msg,err:=tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg fail err=",err)
			return
		}
		//如果读取到消息，下一步处理
		fmt.Printf("msg=%v\n",msg)

	}
}
