package main

import (
	"fmt"
	"net"
	"time"
)

type User struct {
	C    chan string
	ip   string
	name string
}

var users map[string]*User

var message chan string

var exit, hasData chan bool

func main() {
	exit = make(chan bool)
	hasData = make(chan bool)
	message = make(chan string)
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("端口监听失败", err)
		return
	}
	defer listener.Close()

	go managerUsers()

	//首先每一个访问者都需要在独立的协程里去处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("客户端访问失败", err)
			return
		}
		//处理每一个用户的连接
		go handlerFunc(conn)
	}

}

//管理所有的User
func managerUsers() {
	users = make(map[string]*User)

	for {
		msg := <-message
		for _, user := range users {
			user.C <- msg
		}
	}

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
	sendMessage2Users(users[ipStr], "进入聊天室111")


	//接收数据
	go func(){
		for {
			messageBuf := make([]byte, 1024)
			n, err := conn.Read(messageBuf)
			if n == 0 {
				fmt.Println("接收数据出现问题:",err)
				exit <- true
				return
			}
			receiveMsg := string(messageBuf[:n-2])
			fmt.Println("message : ", message)
			if len(receiveMsg) >= 12 && receiveMsg[:11] == "change name" {
				users[ipStr].name = receiveMsg[12:]
				continue
			} else {
				sendMessage2Users(users[ipStr], string(messageBuf[:n-2]))
			}
		}
	}()

	for {
		select {
		case <-exit:
			sendMessage2Users(users[ipStr], "退出聊天室")
			fmt.Println("users = ",users)
			delete(users, ipStr)
			break
		case <-hasData:
			break
		case <-time.After(30 * time.Second):
			break
		}
	}
}

//发送消息
func sendMessage2Users(user *User, content string) {
	message <- fmt.Sprintf("[%s]:%s\n", user.name, content)
}

//用户处理自己内部接收消息的事件
func dealUserMessage(user *User, conn net.Conn) {
	for {
		//堵塞等待消息
		msg := <-user.C
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("用户发送消息失败", err)
			continue
		}
	}
}
