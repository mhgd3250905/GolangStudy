package main

import (
	"fmt"
	"net"
)

type User struct {
	message chan string
	ip      string
	name    string
}

var users map[string]*User

func main() {
	exitChan := make(chan bool)
	users = make(map[string]*User)
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("端口监听失败", err)
		return
	}
	defer listener.Close()

	//首先每一个访问者都需要在独立的协程里去处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("客户端访问失败", err)
			return
		}
		defer conn.Close()
		//处理每一个用户的连接
		go handlerFunc(conn)
	}

	<-exitChan
}

/**
处理每一个用户的连接
 */
func handlerFunc(conn net.Conn) {
	defer conn.Close()
	ipStr := conn.RemoteAddr().String()
	//1.保存用户的信息
	users[ipStr] = &User{make(chan string), ipStr, ipStr}

	//进入每个用户自己的协程
	go dealUserMessage(users[ipStr], conn)

	fmt.Println(conn.RemoteAddr().String(), "连接成功！")
	//2.向每个用户推送连接信息
	for _, user := range users {
		user.message <- fmt.Sprintf("[ %s ] 进入聊天室", users[ipStr].name)
	}

	for {
		messageBuf := make([]byte, 1024)
		n, err := conn.Read(messageBuf)
		if err != nil {
			sendMessage(fmt.Sprintf("[ %s ] 退出聊天室", users[ipStr].name))
			delete(users, ipStr)
			break
		}
		message := string(messageBuf[:n-2])
		fmt.Println("message : ", message)
		if len(message) >= 12 && message[:11] == "change name" {
			users[ipStr].name = message[12:]
			continue
		} else {
			sendMessage(fmt.Sprintf("[ %s ] : %s", users[ipStr].name, string(messageBuf[:n-2])))
		}
	}
}

//发送消息到每一个用户
func sendMessage(content string) {
	for _, user := range users {
		user.message <- content
	}
}

//用户处理自己内部接收消息的事件
func dealUserMessage(user *User, conn net.Conn) {
	for {
		//堵塞等待消息
		message := <-user.message
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("用户发送消息失败", err)
			continue
		}
	}
}