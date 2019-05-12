package main

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	Name string
	Conn net.Conn
}

func main() {
	port := "9999"
	Start(port)
}

var clients map[string]Client

func Start(port string) {
	host := ":" + port

	//获取TCP地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		log.Printf("resolve tcp addr failed:%v\n", err)
		return
	}

	//监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Printf("listen tcp port failed:%v\n", err)
		return
	}

	//消息通道
	msgChan := make(chan string, 10)
	//创建客户端池,用于广播消息
	clients = make(map[string]Client)

	//广播消息
	go BroadMessages(&clients, msgChan)

	//启动
	for {
		fmt.Printf("listening port %s...\n", port)
		//监听有连接接入到服务器
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Accept failed:%v\n", err)
			continue
		}

		//把每个客户端连接扔进连接池
		client := Client{
			Name: conn.RemoteAddr().String(),
			Conn: conn,
		}
		clients[conn.RemoteAddr().String()] = client
		fmt.Println(client, "加入到聊天室")

		//处理消息
		go Handler(conn, &clients, msgChan)

		loginMsg := "系统提示：[" + client.Name + "] 加入到聊天室！"
		msgChan <- loginMsg
	}

}

//处理客户端发送到服务端的消息，并将其扔到通道中
func Handler(conn *net.TCPConn,clients *map[string]Client, msgChan chan string) {
	fmt.Println("connect from client ", conn.RemoteAddr().String())

	buf := make([]byte, 4096)
	for {
		length, err := conn.Read(buf)

		key:=conn.RemoteAddr().String()

		if err != nil {
			log.Printf("read client msg failed:%v\n", err)
			logoutMsg := "系统提示：[" + (*clients)[key].Name + "] 退出聊天室！"
			msgChan <- logoutMsg
			delete(*clients, key)
			conn.Close()
			break
		}

		content := string(buf[:length])
		if len(content) > 8 && content[:7] == "rename|" {
			client := (*clients)[key]
			client.Name=content[7:]
			(*clients)[key]=client
			//把收到的消息写入到通道中
			receiveMsg := "系统提示：[" + (*clients)[key].Name + "]:完成ID修改"
			fmt.Println("change name : ", receiveMsg)
			msgChan <- receiveMsg
		} else {
			//把收到的消息写入到通道中
			receiveMsg := "[" + (*clients)[key].Name + "]:" + string(buf[0:length])
			fmt.Println("read: ", receiveMsg)
			msgChan <- receiveMsg
		}
	}
}

//向所有连接上的客户端发送广播
func BroadMessages(clients *map[string]Client, msgChan chan string) {
	for {
		//不断从通道读取消息
		msg := <-msgChan
		//fmt.Println(msg)

		//向所有客户端发送消息
		for key, client := range *clients {
			fmt.Println("connection is connect from ", key)
			_, err := client.Conn.Write([]byte(msg))
			if err != nil {
				log.Printf("broad msg to %s failed：%v\n", key, err)
				delete(*clients, key)
			}
		}
	}
}
