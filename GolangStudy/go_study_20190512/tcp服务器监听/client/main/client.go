package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("client dial err=", err)
	}

	//功能1：发送单行数据
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标注您输入[终端]

	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string err= ", err)
			return
		}

		line=strings.Trim(line,"\n")
		if line=="exit" {
			fmt.Println("客户端退出...")
			break
		}

		_,err = conn.Write([]byte(line+"\n"))
		if err != nil {
			fmt.Println("conn write err= ", err)
		}
	}

}
