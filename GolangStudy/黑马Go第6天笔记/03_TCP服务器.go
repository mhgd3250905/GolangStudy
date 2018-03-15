package main

import (
	"fmt"
	"net"
	"log"
	"strings"
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
	//函数调用完毕 关闭connect

	fmt.Println(conn.RemoteAddr().String(),"addr connect success!")

	buf:=make([]byte,2048)
	//读取用户数据
	n,err:=conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("buf = \n",string(buf[:n]))

	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
}
