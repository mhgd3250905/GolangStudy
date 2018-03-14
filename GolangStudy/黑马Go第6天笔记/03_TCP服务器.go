package main

import (
	"fmt"
	"net"
	"log"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for{
		conn, err1 := l.Accept()
		if err1 != nil {
			log.Fatal(err)
		}

		//处理用户请求,新建协程
		go HandleRequest(conn)
	}



}

func HandleRequest(conn net.Conn)  {
	fmt.Println(conn.RemoteAddr().String(),"addr connect success!")
	
	//读取用户数据
}
