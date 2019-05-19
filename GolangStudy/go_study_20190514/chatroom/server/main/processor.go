package main

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/server/process2"
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

//先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个serverProcessMsg函数
//功能，根据客户端发送消息种类不同，决定调用那个函数来处理
func (this *Processor)serverProcessMsg(msg *message.Message) (err error) {

	switch msg.Type {
	case message.LoginMsgType:
		//处理登录
		//创建爱你一个UserProcess实例
		up:=&process2.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessLogin(msg)
	case message.RegisterMsgType:
		//创建爱你一个UserProcess实例
		up:=&process2.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessRegister(msg)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}

	return
}

func (this *Processor) Process() (err error) {
	//循环读取客户端消息
	for {

		//这里我么您将读取数据包,直接封装成一个函数readPkg(),返回Message,Err
		//创建一个Transfer完成读包任务
		ty:=&utils.Transfer{
			Conn:this.Conn,
		}
		msg, err := ty.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return err
			} else {
				fmt.Println("readPkg fail err=", err)
				return err
			}
		}

		fmt.Println("msg=", msg)
		err=this.serverProcessMsg(&msg)
		if err != nil {
			fmt.Println("serverProcessMsg fail err=",err)
			return err
		}
	}
}