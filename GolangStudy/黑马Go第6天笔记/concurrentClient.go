package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("访问聊天室失败", err)
		return
	}
	defer conn.Close()

	go readMessage(conn)

	for {
		//读取客户端输入内容
		inputBuf := make([]byte, 1024)
		n, err := os.Stdin.Read(inputBuf)
		if err != nil {
			fmt.Println("读取输入内容失败", err)
			continue
		}
		//fmt.Println("inputBuf: ",string(inputBuf[:n]))
		//发送到服务器
		n, err = conn.Write(inputBuf[:n])
		if err != nil {
			fmt.Println("关闭连接", err)
			break
		}
		//fmt.Println("发送到服务器长度：",n)
	}

}

func readMessage(conn net.Conn) {
	for {
		buf := make([]byte, 2*1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				continue
			}
			break
		}
		fmt.Println(string(buf[:n]))
	}
}