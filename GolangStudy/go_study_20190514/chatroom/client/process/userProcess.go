package process

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/client/utils"
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
	//字段
}

func (this *UserProcess) Register(userId int, userPwd string,
	userName string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var msg message.Message
	msg.Type = message.RegisterMsgType

	//3.创建一个LoginMsg结构体
	var registerMsg message.RegisterMsg
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = userPwd
	registerMsg.User.UserName = userName

	//4.将registerMsg序列化
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给msg.Data字段
	msg.Data = string(data)

	//6.将msg进行序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误err= ", err)
	}

	//这里还需要处理服务器端返回的消息
	msg, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将msg的data部分反序列化为RegisterResMsg
	var registerResMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)
	if registerResMsg.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMsg.Error)
		os.Exit(0)
	}
	return
}

//登陆函数
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	//1.连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var msg message.Message
	msg.Type = message.LoginMsgType

	//3.创建一个LoginMsg结构体
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	//4.将LoginMsg序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5.把data赋给msg.Data字段
	msg.Data = string(data)

	//6.将msg进行序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7.这个时候data就是我们要发送的消息
	//7.1 先把data的长度发送给服务器\
	//先获取到data的长度->转换为一个表示长度的切片

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(bytes) err=", err)
		return
	}

	fmt.Printf("客户端发送消息长度成功,发送消息长度为=%v 内容为=%v\n", len(data), string(data))

	//发送消息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.write(data) fail,err= ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	//这里还需要处理服务器端返回的消息
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//j将msg的data部分反序列化为LoginResMsg
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if loginResMsg.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.User.UserId = userId
		CurUser.User.UserStatus = message.UserOnline

		//fmt.Println("登录成功")

		//现在可以显示一下当前在线用户列表
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMsg.UsersId {
			//不显示自己在线
			if v == userId {
				continue
			}
			fmt.Printf("用户id:%v\t\n", v)
			//完成客户端onlineUsers的初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println()
		fmt.Println()

		//这里还需要在客户端启动一个协程
		//该协程保持和服务器的通讯，如果服务器有数据推送给客户端
		//则接受并显示在客户端终端

		go serverProcessMsg(conn)

		//1.显示登陆成功后的菜单
		ShowMenu()
	} else {
		fmt.Println(loginResMsg.Error)
	}
	return

}
