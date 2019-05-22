package process

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/client/utils"
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"encoding/json"
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
		var content string

		//因为我们总会使用到SmsProcess实例，因为我们将其定义在switch外部
		smsProcess:=&SmsProcess{}

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
			outputOnlineUser()
		case 2:
			fmt.Println("发送消息")
			fmt.Println("你想对大家说点什么:)")
			fmt.Scanf("%s\n",&content)
			smsProcess.sendGroup(content)
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发来的消息")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg fail err=", err)
			return
		}

		//如果读取到消息，下一步处理
		switch msg.Type {
		case message.NotifyUserStatusMsgType: //有人上线了
			//处理
			//1.取出NotifyUserStatusMsg
			var notifyUserStatusMsg message.NotifyUserStatusMsg
			err=json.Unmarshal([]byte(msg.Data),&notifyUserStatusMsg)
			if err != nil {
				fmt.Println("json.Unmarshal fail err=",err)
			}
			//2.把这个用户的状态保存到客户端的map中
			updateUserStatus(&notifyUserStatusMsg)
		case message.SmsMsgType://接收到消息
			//创建一个SmsProcess实例完成转发群聊消息


		default:
			fmt.Println("服务器端返回了一个未知的消息类型...")

		}
		fmt.Printf("msg=%v\n",msg)

	}
}
