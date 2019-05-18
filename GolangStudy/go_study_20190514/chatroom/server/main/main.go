package main

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {


	//不要忘记延时关闭
	defer conn.Close()

	//调用总控，创建一个总控
	processor:=&Processor{
		Conn:conn,
	}

	err:=processor.Process()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程出了问题 fail err= ",err)
		return
	}

}

//编写一个函数，完成对UserDao的初始化任务
func initUserDao(){
	//这里的pool就是一个全局的变量
	model.MyUserDao = model.NewUserDao(pool)
}


func main() {
	//当服务器启东市我们就初始化我么你的redis连接池
	initPool("localhost:6379",16,0,300*time.Second)
	//这里注意初始化顺序
	//初始化UserDao
	initUserDao()

	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.listen err=", err)
	}

	defer listen.Close()

	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来链接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//一旦连接成功则启动一个协程和客户端保持通讯...
		go process(conn)
	}
}
