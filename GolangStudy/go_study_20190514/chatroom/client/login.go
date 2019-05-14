package main

import (
	"../common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//登陆函数
func login(userId int, userPwd string) (err error) {

	//1.连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
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
		return err
	}
	//5.把data赋给msg.Data字段
	msg.Data = string(data)

	//6.将msg进行序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
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
		return err
	}

	fmt.Printf("客户端发送消息长度成功,发送消息长度为=%v 内容为=%v\n",len(data),string(data))
	return nil
}
