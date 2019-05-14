package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	//不要忘记延时关闭
	defer conn.Close()
	//循环读取客户端消息
	for {
		buf := make([]byte, 1024*4)
		fmt.Println("等待读取客户端发送的数据...")
		n,err:=conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err=",err)
			return
		}
		fmt.Println("读到的长度为buf=",buf)
	}
}

func main() {
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
