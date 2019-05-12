package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端发送的数据
	defer conn.Close()//用完必须关闭
	for{
		//创建一个新的切片
		buf:=make([]byte,1024)
		//1.等待客户端通过conn发送信息流
		//2.如果客户端没有write那么这个协程就一直阻塞在这里
		n,err:=conn.Read(buf)//读取数据
		if err == io.EOF {
			fmt.Println("客户端已退出...")
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	//127.0.0.1:8888只有ipv4支持
	//1.tcp表示使用网络协议为tcp
	//2. 0.0.0.0:8888 表示在本地监听8888
	listen,err:=net.Listen("tcp","0.0.0.0:9999")
	if err != nil {
		fmt.Println("listen err= ",err)
		return
	}
	//退出程序关闭接口监听
	defer listen.Close()

	//循环等待客户端来连接我
	for{
		//等待客户端连接
		//fmt.Println("等待客户端连接...")
		conn,err:=listen.Accept()
		if err != nil {
			fmt.Println("Accept() err= ",err)
			continue
		} else {
			fmt.Printf("Accept() suc conn=%v 客户端ip=%v\n",conn,conn.RemoteAddr().String())
		}
		//这里准备一个协程来与客户端进行交互
		go process(conn)
	}

	fmt.Printf("listen suc= %v\n",listen)

}